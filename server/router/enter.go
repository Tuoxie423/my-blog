package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	ArticleRouter
	CommentRouter
	FeedbackRouter
	WebsiteRouter
	HotRouter
	ConfigRouter
	MurmurRouter
}

var RouterGroupApp = new(RouterGroup)
