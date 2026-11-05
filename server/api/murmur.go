package api

import (
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
)

type MurmurApi struct{}

func (m *MurmurApi) MurmurCreate(c *gin.Context) {
	var req request.MurmurCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("The murmur content is invalid.", c)
		return
	}
	var murmur database.Murmur
	murmur.Content = req.Content
	err := murmurService.MurmurCreate(murmur)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("发布碎碎念失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("发布碎碎念成功", c)

}

func (m *MurmurApi) MurmurDelete(c *gin.Context) {
	var ids request.MurmurDelete
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}
	if len(ids.Ids) == 0 {
		response.FailWithMessage("No ids provided", c)
		return
	}
	err := murmurService.MurmurDelete(ids)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("删除碎碎念失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("删除碎碎念成功", c)

}

func (m *MurmurApi) MurmurList(c *gin.Context) {
	list, err := murmurService.MurmurList()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("加载碎碎念失败，请稍后重试", c)
		return
	}
	response.OkWithData(list, c)
}
