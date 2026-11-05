package api

import (
	"server/global"
	"server/model/request"
	"server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageApi struct {
}

// ImageUpload 上传图片
func (imageApi *ImageApi) ImageUpload(c *gin.Context) {
	_, header, err := c.Request.FormFile("image")
	if err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		response.FailWithMessage("No image was uploaded, or the file is invalid.", c)
		return
	}

	url, err := imageService.ImageUpload(header)
	if err != nil {
		global.Log.Error("Failed to upload image:", zap.Error(err))
		response.FailWithMessage("上传图片失败，请稍后重试", c)
		return
	}
	// local 返回格式 /uploads/image/fileName | qiniu 返回格式 http(s)://image.xxx.xx/fileName
	response.OkWithDetailed(response.ImageUpload{
		Url:     url,
		OssType: global.Config.System.OssType,
	}, "Successfully uploaded image", c)
}

// ImageDelete 删除图片
func (imageApi *ImageApi) ImageDelete(c *gin.Context) {
	var req request.ImageDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}

	err = imageService.ImageDelete(req)
	if err != nil {
		global.Log.Error("Failed to delete image:", zap.Error(err))
		response.FailWithMessage("删除图片失败，请稍后重试", c)
		return
	}
	response.OkWithMessage("删除图片成功", c)
}

// ImageList 获取图片列表
func (imageApi *ImageApi) ImageList(c *gin.Context) {
	var pageInfo request.ImageList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("The request parameters are invalid.", c)
		return
	}

	imageList, total, err := imageService.ImageList(pageInfo)
	if err != nil {
		global.Log.Error("Failed to get image list:", zap.Error(err))
		response.FailWithMessage("获取图片列表失败，请稍后重试", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  imageList,
		Total: total,
	}, c)
}
