package router

import (
	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

func (c *CommentRouter) InitCommentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
}
