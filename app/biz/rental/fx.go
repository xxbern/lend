package rental

import (
	"go.uber.org/fx"
	"lend/app/biz/rental/api"
	"lend/gen/oas"
)

var Module = fx.Options(
	fx.Provide(api.NewLendHandler),
	fx.Provide(
		fx.Annotate(
			api.NewLendHandler,
			fx.As(new(oas.LendBizHandler)),
		),
	),
)
