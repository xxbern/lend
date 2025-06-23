package server

import (
	"context"
	"go.uber.org/fx"
	"net/http"
)

// Module 注入 Fx 的模块
var Module = fx.Options(
	fx.Provide(NewServerHandler),
	fx.Provide(NewHTTPServer),
	fx.Invoke(func(lc fx.Lifecycle, server *http.Server) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return OnStart(server)
			},
			OnStop: func(ctx context.Context) error {
				return OnStop(ctx, server)
			},
		})
	}),
)
