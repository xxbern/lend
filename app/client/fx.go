package client

import (
	"go.uber.org/fx"
	"lend/app/client/wx"
)

var Module = fx.Options(
	fx.Provide(wx.Init),
)
