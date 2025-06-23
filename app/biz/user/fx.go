package user

import (
	"go.uber.org/fx"
	"lend/app/biz/user/api"
	"lend/app/biz/user/dba"
	"lend/gen/oas"
)

var UserModule = fx.Options(
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
)
