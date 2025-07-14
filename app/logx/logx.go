package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	Logger *zap.Logger
)

func Init(production bool) {
	// 1. 定义 Encoder（控制日志格式）
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime) // 时间格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 2. 定义 stdout（INFO）和 stderr（WARN/ERROR）
	stdout := zapcore.AddSync(os.Stdout)
	stderr := zapcore.AddSync(os.Stderr)

	// 3. 根据 production 模式设置 info 日志的最低级别
	var infoLevel zap.LevelEnablerFunc
	// 5. 构建 Logger
	var opts []zap.Option
	if production {
		infoLevel = func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel // 仅输出 info
		}
	} else {
		infoLevel = func(lvl zapcore.Level) bool {
			return lvl <= zapcore.InfoLevel // 输出 debug 和 info
		}
		opts = append(opts, zap.AddCaller()) // 输出代码位置
	}

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel // WARN/ERROR → stderr
	})

	// 4. 创建 Core
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, stdout, infoLevel),  // info/debug（取决于模式）→ stdout
		zapcore.NewCore(encoder, stderr, errorLevel), // warn/error → stderr
	)

	Logger = zap.New(core, opts...)

	defer Logger.Sync()
	RedirectStdLog()
}
