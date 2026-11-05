package api

import (
	"server/config"
	"server/global"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConfigApi struct {
}

// GetWebsite 获取网站配置
func (configApi *ConfigApi) GetWebsite(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

// UpdateWebsite 更新网站配置
func (configApi *ConfigApi) UpdateWebsite(c *gin.Context) {
	var req config.Website
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The website configuration is invalid.", c)
		return
	}

	err = configService.UpdateWebsite(req)
	if err != nil {
		global.Log.Error("Failed to update website:", zap.Error(err))
		response.FailWithMessage("更新网站配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新网站配置成功", c)
}

// GetSystem 获取系统配置
func (configApi *ConfigApi) GetSystem(c *gin.Context) {
	response.OkWithData(global.Config.System, c)
}

// UpdateSystem 更新系统配置
func (configApi *ConfigApi) UpdateSystem(c *gin.Context) {
	var req config.System
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The system configuration is invalid.", c)
		return
	}

	err = configService.UpdateSystem(req)
	if err != nil {
		global.Log.Error("Failed to update system:", zap.Error(err))
		response.FailWithMessage("更新系统配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新系统配置成功", c)
}

// GetEmail 获取邮箱配置
func (configApi *ConfigApi) GetEmail(c *gin.Context) {
	response.OkWithData(global.Config.Email, c)
}

// UpdateEmail 更新邮箱配置
func (configApi *ConfigApi) UpdateEmail(c *gin.Context) {
	var req config.Email
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The email configuration is invalid.", c)
		return
	}

	err = configService.UpdateEmail(req)
	if err != nil {
		global.Log.Error("Failed to update email:", zap.Error(err))
		response.FailWithMessage("更新邮箱配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新邮箱配置成功", c)
}

// GetQQ 获取QQ登录配置
func (configApi *ConfigApi) GetQQ(c *gin.Context) {
	response.OkWithData(global.Config.QQ, c)
}

// UpdateQQ 更新QQ登录配置
func (configApi *ConfigApi) UpdateQQ(c *gin.Context) {
	var req config.QQ
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The QQ configuration is invalid.", c)
		return
	}

	err = configService.UpdateQQ(req)
	if err != nil {
		global.Log.Error("Failed to update qq:", zap.Error(err))
		response.FailWithMessage("更新 QQ 配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新 QQ 配置成功", c)
}

// GetQiniu 获取七牛云配置
func (configApi *ConfigApi) GetQiniu(c *gin.Context) {
	response.OkWithData(global.Config.Qiniu, c)
}

// UpdateQiniu 更新七牛云配置
func (configApi *ConfigApi) UpdateQiniu(c *gin.Context) {
	var req config.Qiniu
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The Qiniu configuration is invalid.", c)
		return
	}

	err = configService.UpdateQiniu(req)
	if err != nil {
		global.Log.Error("Failed to update qiniu:", zap.Error(err))
		response.FailWithMessage("更新七牛配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新七牛配置成功", c)
}

// GetJwt 获取Jwt配置
func (configApi *ConfigApi) GetJwt(c *gin.Context) {
	response.OkWithData(global.Config.Jwt, c)
}

// UpdateJwt 更新Jwt配置
func (configApi *ConfigApi) UpdateJwt(c *gin.Context) {
	var req config.Jwt
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The JWT configuration is invalid.", c)
		return
	}

	err = configService.UpdateJwt(req)
	if err != nil {
		global.Log.Error("Failed to update jwt:", zap.Error(err))
		response.FailWithMessage("更新 JWT 配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新 JWT 配置成功", c)
}

// GetGaode 获取高德配置
func (configApi *ConfigApi) GetGaode(c *gin.Context) {
	response.OkWithData(global.Config.Gaode, c)
}

// UpdateGaode 更新高德配置
func (configApi *ConfigApi) UpdateGaode(c *gin.Context) {
	var req config.Gaode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The Gaode configuration is invalid.", c)
		return
	}

	err = configService.UpdateGaode(req)
	if err != nil {
		global.Log.Error("Failed to update gaode:", zap.Error(err))
		response.FailWithMessage("更新高德配置失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("更新高德配置成功", c)
}
