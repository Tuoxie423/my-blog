package service

type ServiceGroup struct {
	EsService
	BaseService
	UserService
	JwtService
	GaodeService
	QQService
	ImageService
	ArticleService
	CommentService
	FeedbackService
}

var ServiceGroupApp = new(ServiceGroup)
