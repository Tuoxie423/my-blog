package api

import (
	"errors"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// MurmurLike 切换赞同状态
func (m *MurmurApi) MurmurLike(c *gin.Context) {
	var req request.MurmurLike
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}

	req.UserID = utils.GetUserID(c)
	if req.UserID == 0 {
		response.NoAuth("请先登录", c)
		return
	}

	result, err := murmurService.MurmurLike(req)
	if err != nil {
		global.Log.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("碎碎念不存在或已删除", c)
			return
		}
		response.FailWithMessage("操作失败，请稍后重试", c)
		return
	}
	response.OkWithData(result, c)
}

// MurmurLiked 当前用户点过赞的碎碎念 ID 列表
func (m *MurmurApi) MurmurLiked(c *gin.Context) {
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.NoAuth("请先登录", c)
		return
	}

	ids, err := murmurService.MurmurLiked(userID)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("加载赞同状态失败，请稍后重试", c)
		return
	}
	response.OkWithData(ids, c)
}
