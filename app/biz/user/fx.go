package user

import (
	"go.uber.org/fx"
	"lend/app/biz/user/api"
	"lend/app/biz/user/internal"
	"lend/app/biz/user/internal/dba"
	"lend/app/biz/user/service"
	"lend/gen/oas"
)

// fx作为依赖注入框架，只关注实例构建和依赖管理，不应被其他代码依赖
// 如果移除项目内所有fx文件，手动创建所有实例也可以保证程序运行，不要对fx进行强依赖

var Module = fx.Options(
	fx.Provide(dba.NewUserRepository),
	fx.Provide(
		fx.Annotate(
			api.NewUserManageHandler,
			// As 声明为oas相关包的实现 server支持注入oas接口
			fx.As(new(oas.UserManageHandler)),
		),
	),
	fx.Provide(
		fx.Annotate(
			api.NewUserInfoHandler,
			fx.As(new(oas.UserInfoHandler)),
		),
	),
	fx.Provide(
		fx.Annotate(
			api.NewWxLoginHandler,
			fx.As(new(oas.WxLoginHandler)),
		),
	),
	fx.Provide(
		fx.Annotate(
			internal.NewUserServiceImpl,
			// 公开功能
			fx.As(new(service.PubUserService)),
			// 内部功能
			fx.As(new(internal.UserServiceInternal)),
		)),
)
