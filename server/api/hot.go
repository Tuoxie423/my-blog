package api

import (
	"server/model/response"

	"github.com/gin-gonic/gin"
)

type HotApi struct{}

func (h *HotApi) GetHot(c *gin.Context) {
	list, err := hotService.GetHotAll()
	if err != nil {
		response.FailWithMessage("加载热榜失败，请稍后重试", c)
		return
	}
	response.OkWithData(list, c)
}
