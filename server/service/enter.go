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
	MurmurService
}

var ServiceGroupApp = new(ServiceGroup)
