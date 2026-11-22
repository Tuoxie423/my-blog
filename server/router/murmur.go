package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type MurmurRouter struct{}

func (m *MurmurRouter) InitMurmurRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	murmurAdminRouter := Router.Group("murmur")
	murmurPublicRouter := PublicRouter.Group("murmur")
	murmurPrivateRouter := PrivateRouter.Group("murmur")
	murmurApi := api.ApiGroupApp.MurmurApi
	{
		murmurAdminRouter.POST("create", murmurApi.MurmurCreate)
		murmurAdminRouter.DELETE("delete", murmurApi.MurmurDelete)

	}
	{
		// 登录用户即可点赞，不需要管理员权限
		murmurPrivateRouter.POST("like", murmurApi.MurmurLike)
		murmurPrivateRouter.GET("liked", murmurApi.MurmurLiked)
	}
	{
		murmurPublicRouter.GET("all", murmurApi.MurmurList)
	}

}
