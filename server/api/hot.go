package api

import (
	"server/model/response"

	"github.com/gin-gonic/gin"
)

type HotApi struct{}

func (h *HotApi) GetHot(c *gin.Context) {
	list, err := hotService.GetHotAll()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
