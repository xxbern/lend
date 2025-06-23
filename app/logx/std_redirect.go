package logx

import (
	"go.uber.org/zap"
	"log"
)

func RedirectStdLog() {
	log.SetFlags(0)
	log.SetOutput(zap.NewStdLog(Logger).Writer())
}
