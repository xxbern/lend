package biz

import (
	"go.uber.org/fx"
	"lend/app/biz/rental"
	"lend/app/biz/user"
)

var Module = fx.Options(
	user.UserModule,
	rental.RentalModule,
)
