package db

import (
	"context"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"lend/app/config"
	"lend/gen/db"
)

// 注入 Fx 的模块
var Module = fx.Options(
	fx.Provide(NewDB, db.Use),
	fx.Invoke(func(lc fx.Lifecycle, db *gorm.DB, cfg *config.Config) {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return OnStart(ctx, db, cfg)
			},
			OnStop: func(ctx context.Context) error {
				return OnStop(ctx, db, cfg)
			},
		})
	},
	),
)
