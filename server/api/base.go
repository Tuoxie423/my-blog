package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

// Captcha 生成数字验证码
func (baseApi *BaseApi) Captcha(c *gin.Context) {
	// 创建数字验证码的驱动
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	// 创建验证码对象
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	id, b64s, _, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha:", zap.Error(err))
		response.FailWithMessage("生成验证码失败，请稍后重试", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

// SendEmailVerificationCode 发送邮箱验证码
func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The email address or verification code is invalid.", c)
		return
	}
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		err = baseService.SendEmailVerificationCode(c, req.Email)
		if err != nil {
			global.Log.Error("Failed to send email:", zap.Error(err))
			response.FailWithMessage("发送邮件失败，请稍后重试", c)
			return
		}
		response.OkWithMessage("邮件发送成功", c)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

// QQLoginURL 返回 QQ 登录链接
func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}
