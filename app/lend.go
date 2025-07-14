package main

import (
	"go.uber.org/fx"
	"lend/app/biz"
	"lend/app/client"
	"lend/app/config"
	"lend/app/db"
	"lend/app/logx"
	"lend/app/server"
)

func main() {
	logx.Init(true)

	app := fx.New(
		logx.FxLogger,
		config.Module,
		db.Module,
		server.Module,
		client.Module,
		biz.Module,
	)

	app.Run()
}
