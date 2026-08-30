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
		response.FailWithMessage(err.Error(), c)
		return
	}
	var murmur database.Murmur
	murmur.Content = req.Content
	err := murmurService.MurmurCreate(murmur)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Successfully created a new murmur! ", c)

}

func (m *MurmurApi) MurmurDelete(c *gin.Context) {
	var ids request.MurmurDelete
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(ids.Ids) == 0 {
		response.FailWithMessage("No ids provided", c)
		return
	}
	err := murmurService.MurmurDelete(ids)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Successfully deleted the murmur! ", c)

}

func (m *MurmurApi) MurmurList(c *gin.Context) {
	list, err := murmurService.MurmurList()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
