package logx

import (
	"fmt"
	"go.uber.org/fx"
)

var FxLogger = fx.Logger(FuncPrinter{})

type FuncPrinter struct {
	f func(format string, args ...interface{})
}

func (FuncPrinter) Printf(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...))
}
