package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type WebsiteRouter struct {
}

func (w *WebsiteRouter) InitWebsiteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	websiteRouter := Router.Group("website")
	websitePublicRouter := PublicRouter.Group("website")

	websiteApi := api.ApiGroupApp.WebsiteApi
	{
		websiteRouter.POST("createFooterLink", websiteApi.WebsiteCreateFooterLink)
		websiteRouter.DELETE("deleteFooterLink", websiteApi.WebsiteDeleteFooterLink)
	}
	{
		websitePublicRouter.GET("logo", websiteApi.WebsiteLogo)
		websitePublicRouter.GET("title", websiteApi.WebsiteTitle)
		websitePublicRouter.GET("info", websiteApi.WebsiteInfo)
		websitePublicRouter.GET("footerLink", websiteApi.WebsiteFooterLink)
		websitePublicRouter.GET("yiyan", websiteApi.WebsiteYiyan)
	}
}
