package config

import "go.uber.org/fx"

// 注入 Fx 的模块
var Module = fx.Options(
	fx.Provide(Init),
)
