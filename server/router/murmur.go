package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type MurmurRouter struct{}

func (m *MurmurRouter) InitMurmurRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	murmurAdminRouter := Router.Group("murmur")
	murmurPublicRouter := PublicRouter.Group("murmur")
	murmurApi := api.ApiGroupApp.MurmurApi
	{
		murmurAdminRouter.POST("create", murmurApi.MurmurCreate)
		murmurAdminRouter.DELETE("delete", murmurApi.MurmurDelete)

	}
	{
		murmurPublicRouter.GET("all", murmurApi.MurmurList)
	}

}
