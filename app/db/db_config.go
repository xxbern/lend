package db

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lend/app/config"
	"lend/app/logx"
	"lend/gen/db"
	"regexp"
	"time"
)

//
//var (
//	Db    *gorm.DB
//	query *db.Query
//)

// 注入 Fx 的模块
var Module = fx.Options(
	fx.Provide(NewDB, db.Use),
	fx.Invoke(RegisterDBHooks),
)

// NewDB 创建 *gorm.DB 连接（不直接打开连接，由生命周期管理）
func NewDB(cfg *config.Config) (*gorm.DB, error) {
	connUrl := cfg.DB.ConnUrl()

	// 掩码化连接字符串中的密码（安全日志）
	re := regexp.MustCompile(`(?i)(password\s*=\s*)[^ ]+`)
	maskedConnUrl := re.ReplaceAllString(connUrl, "${1}******")
	logx.Logger.Info("Initializing database connection", zap.String("url", maskedConnUrl))

	// 初始化 GORM（不立即连接）
	gdb, err := gorm.Open(postgres.Open(connUrl), &gorm.Config{
		Logger:                 logx.NewGormLogger(logger.Info, 200*time.Millisecond),
		SkipDefaultTransaction: true, // 可选：提升性能
		PrepareStmt:            true, // 可选：预编译 SQL
	})
	if err != nil {
		logx.Logger.Error("Failed to initialize DB", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize DB: %w", err)
	}

	return gdb, nil
}

// RegisterDBHooks 注册数据库生命周期钩子
func RegisterDBHooks(lc fx.Lifecycle, db *gorm.DB, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 获取底层 SQL.DB 设置连接池
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get SQL.DB: %w", err)
			}

			// 配置连接池（根据实际需求调整）
			sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConn) // 默认值建议 10-100
			sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConn) // 默认值建议 2-10
			sqlDB.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime)
			sqlDB.SetConnMaxIdleTime(cfg.DB.ConnMaxLifetime)
			// 测试连接是否可用
			if err := sqlDB.PingContext(ctx); err != nil {
				logx.Logger.Error("Failed to ping DB", zap.Error(err))
				return fmt.Errorf("DB ping failed: %w", err)
			}

			logx.Logger.Info("Database connection established")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 关闭数据库连接
			sqlDB, err := db.DB()
			if err != nil {
				logx.Logger.Error("Failed to get SQL.DB on shutdown", zap.Error(err))
				return nil // 即使失败也继续关闭流程
			}

			if err := sqlDB.Close(); err != nil {
				logx.Logger.Error("Failed to close DB", zap.Error(err))
				return fmt.Errorf("failed to close DB: %w", err)
			}

			logx.Logger.Info("Database connection closed")
			return nil
		},
	})
}
