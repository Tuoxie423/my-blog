package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	ArticleRouter
	CommentRouter
	FeedbackRouter
}

var RouterGroupApp = new(RouterGroup)
