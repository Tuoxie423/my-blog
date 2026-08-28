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
	WebsiteService
	HotService
	ConfigService
}

var ServiceGroupApp = new(ServiceGroup)
