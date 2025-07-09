package api

import (
	"go.uber.org/fx"
	"lend/gen/oas"
)

// fx作为依赖注入框架，只关注实例构建和依赖管理，不应被其他代码依赖
// 如果移除项目内所有fx文件，手动创建所有实例也可以保证程序运行，不要对fx进行强依赖

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewUserManageHandler,
			// As 声明为oas相关包的实现 server支持注入oas接口
			fx.As(new(oas.UserManageHandler)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewUserInfoHandler,
			fx.As(new(oas.UserInfoHandler)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewWxLoginHandler,
			fx.As(new(oas.WxLoginHandler)),
		),
	),
)
