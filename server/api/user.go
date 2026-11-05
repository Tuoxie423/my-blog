package api

import (
	"errors"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type UserApi struct{}

func (userApi *UserApi) Register(c *gin.Context) {
	var req request.Register
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The registration information is invalid.", c)
		return
	}
	session := sessions.Default(c)
	// 两次邮箱和验证码的一致性判断
	savedEmail := session.Get("email")
	if savedEmail == nil || savedEmail.(string) != req.Email {
		response.FailWithMessage("邮箱与待验证的邮箱不一致", c)
		return
	}
	savedCode := session.Get("verification_code")
	if savedCode == nil || savedCode.(string) != req.VerificationCode {
		response.FailWithMessage("验证码错误", c)
		return
	}

	savedTime := session.Get("expire_time")
	if savedTime == nil || savedTime.(int64) < time.Now().Unix() {
		response.FailWithMessage("验证码已过期，请重新发送", c)
		return
	}

	u := database.User{Username: req.Username, Password: req.Password, Email: req.Email}

	user, err := userService.Register(u)
	if err != nil {
		global.Log.Error("Failed to register user:", zap.Error(err))
		response.FailWithMessage("注册失败，请稍后重试", c)
		return
	}

	// 注册成功后，生成 token 并返回
	userApi.TokenNext(c, user)

}

func (userApi *UserApi) Login(c *gin.Context) {
	switch c.Query("flag") {
	case "email":
		userApi.EmailLogin(c)
	case "qq":
		userApi.QQLogin(c)
	default:
		userApi.EmailLogin(c)
	}
}

func (userApi *UserApi) EmailLogin(c *gin.Context) {
	var req request.Login
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The login information is invalid.", c)
		return
	}
	if store.Verify(req.CaptchaId, req.Captcha, true) {
		u := database.User{Email: req.Email, Password: req.Password}
		user, err := userService.EmailLogin(u)
		if err != nil {
			global.Log.Error("Failed to login user:", zap.Error(err))
			response.FailWithMessage("登录失败，请稍后重试", c)
			return
		}

		userApi.TokenNext(c, user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

func (userApi *UserApi) QQLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("缺少 code 参数", c)
		return
	}
	accessTokenResponse, err := qqService.GetAccessTokenByCode(code)
	if err != nil {
		response.FailWithMessage("QQ 授权失败，请重试", c)
		return
	}
	user, err := userService.QQLogin(accessTokenResponse)
	if err != nil {
		global.Log.Error("Failed to login user:", zap.Error(err))
		response.FailWithMessage("登录失败，请稍后重试", c)
		return
	}
	userApi.TokenNext(c, user)
}

func (userApi *UserApi) TokenNext(c *gin.Context, user database.User) {
	if user.Freeze {
		response.FailWithMessage("账号已被冻结，请联系管理员", c)
		return
	}

	baseClaims := request.BaseClaims{
		UserID: user.ID,
		UUID:   user.UUID,
		RoleID: user.RoleID,
	}

	j := utils.NewJWT()

	accessClaims := j.CreateAccessClaims(baseClaims)
	accessToken, err := j.CreateAccessToken(accessClaims)
	if err != nil {
		global.Log.Error("Failed to create access token:", zap.Error(err))
		response.FailWithMessage("登录凭证生成失败，请重新登录", c)
		return
	}

	refreshClaims := j.CreateRefreshClaims(baseClaims)
	refreshToken, err := j.CreateRefreshToken(refreshClaims)
	if err != nil {
		global.Log.Error("Failed to create refresh token:", zap.Error(err))
		response.FailWithMessage("登录凭证生成失败，请重新登录", c)
		return
	}

	if !global.Config.System.UseMultipoint {
		// 设置刷新令牌并返回
		utils.SetRefreshToken(c, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "Successful login", c)
		return
	}
	if jwtStr, err := jwtService.GetRedisJWT(user.UUID); errors.Is(err, redis.Nil) {
		// 不存在就设置新的
		if err := jwtService.SetRedisJWT(refreshToken, user.UUID); err != nil {
			global.Log.Error("Failed to set login status:", zap.Error(err))
			response.FailWithMessage("登录状态设置失败，请稍后重试", c)
			return
		}
		utils.SetRefreshToken(c, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "Successful login", c)
		return
	} else {
		var blacklist database.JwtBlacklist
		blacklist.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blacklist); err != nil {
			global.Log.Error("Failed to invalidate jwt:", zap.Error(err))
			response.FailWithMessage("登录凭证作废失败，请稍后重试", c)
			return
		}

		// 设置新的 JWT 到 Redis
		if err := jwtService.SetRedisJWT(refreshToken, user.UUID); err != nil {
			global.Log.Error("Failed to set login status:", zap.Error(err))
			response.FailWithMessage("登录状态设置失败，请稍后重试", c)
			return
		}

		// 设置刷新令牌并返回
		utils.SetRefreshToken(c, refreshToken, int(refreshClaims.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                 user,
			AccessToken:          accessToken,
			AccessTokenExpiresAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "Successful login", c)
		return
	}
}

func (userApi *UserApi) ForgotPassword(c *gin.Context) {
	var req request.ForgotPassword
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}
	session := sessions.Default(c)
	// 两次邮箱和验证码的一致性判断
	savedEmail := session.Get("email")
	if savedEmail == nil || savedEmail.(string) != req.Email {
		response.FailWithMessage("邮箱与待验证的邮箱不一致", c)
		return
	}
	savedCode := session.Get("verification_code")
	if savedCode == nil || savedCode.(string) != req.VerificationCode {
		response.FailWithMessage("验证码错误", c)
		return
	}

	savedTime := session.Get("expire_time")
	if savedTime == nil || savedTime.(int64) < time.Now().Unix() {
		response.FailWithMessage("验证码已过期，请重新发送", c)
		return
	}
	err = userService.ForgotPassword(req)
	if err != nil {
		global.Log.Error("Failed to reset password:", zap.Error(err))
		response.FailWithMessage("密码重置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("密码重置成功", c)
}

func (userApi *UserApi) UserCard(c *gin.Context) {
	var req request.UserCard
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage("The user ID is invalid.", c)
		return
	}
	userCard, err := userService.UserCard(req)
	if err != nil {
		global.Log.Error("Failed to get user card:", zap.Error(err))
		response.FailWithMessage("获取用户卡片失败，请稍后重试", c)
		return
	}
	response.OkWithData(userCard, c)
}

func (userApi *UserApi) Logout(c *gin.Context) {
	userService.Logout(c)
	response.OkWithMessage("退出登录成功", c)
}

func (userApi *UserApi) UserResetPassword(c *gin.Context) {
	var req request.UserResetPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("The password information is invalid.", c)
		return
	}
	req.UserID = utils.GetUserID(c)
	err := userService.UserResetPassword(req)
	if err != nil {
		global.Log.Error("Failed to reset password:", zap.Error(err))
		response.FailWithMessage("密码重置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("密码重置成功", c)
}

func (userApi *UserApi) UserInfo(c *gin.Context) {
	userID := utils.GetUserID(c)
	user, err := userService.UserInfo(userID)
	if err != nil {
		global.Log.Error("Failed to get user info:", zap.Error(err))
		response.FailWithMessage("获取用户信息失败，请稍后重试", c)
		return
	}
	response.OkWithData(user, c)
}

// UserChangeInfo 修改个人信息
func (userApi *UserApi) UserChangeInfo(c *gin.Context) {
	var req request.UserChangeInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The profile information is invalid.", c)
		return
	}
	req.UserID = utils.GetUserID(c)
	err = userService.UserChangeInfo(req)
	if err != nil {
		global.Log.Error("Failed to change user information:", zap.Error(err))
		response.FailWithMessage("修改个人信息失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("修改个人信息成功", c)
}

func (userApi *UserApi) UserWeather(c *gin.Context) {
	ip := c.ClientIP()
	weather, err := userService.UserWeather(ip)
	if err != nil {
		global.Log.Error("获取天气失败，请稍后重试", zap.Error(err))
		response.FailWithMessage("获取天气失败，请稍后重试", c)
		return
	}
	response.OkWithData(weather, c)
}

func (userApi *UserApi) UserList(c *gin.Context) {
	var pageInfo request.UserList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}

	list, total, err := userService.UserList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get user list:", zap.Error(err))
		response.FailWithMessage("获取用户列表失败，请稍后重试", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

func (userApi *UserApi) UserFreeze(c *gin.Context) {
	var req request.UserOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}
	err = userService.UserFreeze(req)
	if err != nil {
		global.Log.Error("Failed to freeze user:", zap.Error(err))
		response.FailWithMessage("冻结用户失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("冻结用户成功", c)
}

func (userApi *UserApi) UserUnfreeze(c *gin.Context) {
	var req request.UserOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}
	err = userService.UserUnfreeze(req)
	if err != nil {
		global.Log.Error("Failed to unfreeze user:", zap.Error(err))
		response.FailWithMessage("解冻用户失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("解冻用户成功", c)
}

func (userApi *UserApi) UserLoginList(c *gin.Context) {
	var pageInfo request.UserLoginList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}

	list, total, err := userService.UserLoginList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get user login list:", zap.Error(err))
		response.FailWithMessage("获取登录记录失败，请稍后重试", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
