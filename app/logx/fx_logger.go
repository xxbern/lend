package logx

import (
	"fmt"
	"go.uber.org/fx"
)

var FxLogger = fx.Logger(FuncPrinter{})

type FuncPrinter struct {
}

func (FuncPrinter) Printf(format string, args ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, args...))
}
