package api

import (
	"net/http"
	"server/global"
	"server/model/database"
	"server/model/response"

	uapi "github.com/AxT-Team/uapi-sdk-go/uapi"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebsiteApi struct {
}

// WebsiteLogo 网站 Logo 链接
func (website *WebsiteApi) WebsiteLogo(c *gin.Context) {
	if global.Config.Website.Logo != "" {
		c.Redirect(http.StatusMovedPermanently, global.Config.Website.Logo)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/image/logo.png")
	}
}

// WebsiteTitle 网站标题栏
func (website *WebsiteApi) WebsiteTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"title": global.Config.Website.Title})
}

// WebsiteInfo 获取网站信息
func (website *WebsiteApi) WebsiteInfo(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

// WebsiteFooterLink 获取页脚链接
func (website *WebsiteApi) WebsiteFooterLink(c *gin.Context) {
	footerLinks := websiteService.WebsiteFooterLink()
	response.OkWithData(footerLinks, c)
}

// WebsiteCreateFooterLink 创建页脚链接
func (website *WebsiteApi) WebsiteCreateFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteCreateFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to create footer link:", zap.Error(err))
		response.FailWithMessage("Failed to create footer link", c)
		return
	}
	response.OkWithMessage("Successfully created footer link", c)
}

// WebsiteDeleteFooterLink 删除页脚链接
func (website *WebsiteApi) WebsiteDeleteFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.WebsiteDeleteFooterLink(req)

	if err != nil {
		global.Log.Error("Failed to delete footer link:", zap.Error(err))
		response.FailWithMessage("Failed to delete footer link", c)
		return
	}
	response.OkWithMessage("Successfully deleted footer link", c)
}

func (website *WebsiteApi) WebsiteYiyan(c *gin.Context) {
	if global.Config.Yiyan.Enabled {
		client := uapi.New("https://uapis.cn", "")
		params := map[string]any{}
		resp, err := client.Poem().GetSaying(params)
		if err != nil {
			global.Log.Error("Failed to call API(uapis.cn):", zap.Error(err))
			response.OkWithData(global.Config.Yiyan.Default, c)
			return
		}
		saying := global.Config.Yiyan.Default
		if m, ok := resp.(map[string]any); ok {
			if s, ok := m["text"].(string); ok && s != "" {
				saying = s
			}
		}
		response.OkWithData(saying, c)
	} else {
		response.OkWithData(global.Config.Yiyan.Default, c)
	}
}
