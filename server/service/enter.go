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
}

var ServiceGroupApp = new(ServiceGroup)
