package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func Init(production bool) {

	// 1. 定义 Encoder（控制日志格式）
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 2. 定义 stdout（INFO/DEBUG）和 stderr（WARN/ERROR）
	stdout := zapcore.AddSync(os.Stdout)
	stderr := zapcore.AddSync(os.Stderr)

	// 3. 定义日志级别分流规则
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel // DEBUG/INFO → stdout
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel // WARN/ERROR → stderr
	})

	// 4. 创建 Core（合并 stdout 和 stderr）
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, stdout, infoLevel),  // INFO/DEBUG → stdout
		zapcore.NewCore(encoder, stderr, errorLevel), // WARN/ERROR → stderr
	)

	// 5. 根据 production 模式添加额外配置
	var opts []zap.Option
	if !production {
		opts = append(opts, zap.AddCaller()) // 开发模式显示调用位置
	}

	// 6. 创建 Logger
	Logger = zap.New(core, opts...)
	Sugar = Logger.Sugar()

	defer Logger.Sync()
	RedirectStdLog()
}
