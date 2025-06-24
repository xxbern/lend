package server

import (
	"context"
	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"lend/app/config"
	"lend/app/logx"
	"lend/gen/oas"
	"net/http"
	"time"
)

// NewHTTPServer 构造 handler 和 http.Server
func NewHTTPServer(cfg *config.Config, handler oas.Handler) (*http.Server, error) {
	oasServer, err := oas.NewServer(
		handler,
		SecurityHandler{},
		oas.WithMiddleware(logging()),
		oas.WithErrorHandler(errorHandle()),
	)
	if err != nil {
		logx.Logger.Error("Failed to create server", zap.Error(err))
		return nil, err
	}

	addr := cfg.Server.Addr()

	return &http.Server{
		Addr:              addr,
		Handler:           oasServer,
		ReadHeaderTimeout: 5 * time.Second,  // 增加超时时间
		IdleTimeout:       30 * time.Second, // 新增空闲连接超时
	}, nil
}

func OnStart(server *http.Server) error {
	go func() {
		logx.Logger.Info("HTTP server started", zap.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logx.Logger.Error("Server stopped with error", zap.Error(err))
		}
	}()
	return nil
}

func OnStop(ctx context.Context, server *http.Server) error {
	logx.Logger.Info("Shutting down HTTP server...")
	return server.Shutdown(ctx)
}
