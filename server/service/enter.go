package service

type ServiceGroup struct {
	EsService
	BaseService
	UserService
	JwtService
	GaodeService
	QQService
}

var ServiceGroupApp = new(ServiceGroup)
