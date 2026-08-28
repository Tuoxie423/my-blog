package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type HotRouter struct {
}

func (h *HotRouter) InitHotRouter(PublicRouter *gin.RouterGroup) {
	hotPublicRouter := PublicRouter.Group("hot")

	HotApi := api.ApiGroupApp.HotApi
	hotPublicRouter.GET("all", HotApi.GetHot)
}
