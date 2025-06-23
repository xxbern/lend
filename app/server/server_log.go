package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/middleware"
	"go.uber.org/zap"
	"lend/app/logx"
	"time"
)

func logging() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		reqId := req.Raw.Header.Get("X-Request-ID")
		url := req.Raw.URL
		if reqId == "" {
			reqId = uuid.New().String()
		}
		c := context.WithValue(req.Context, "requestID", reqId)
		req.SetContext(c)

		// 记录请求开始
		startTime := time.Now()
		logx.Logger.Info("Handling request",
			zap.String("requestID", reqId),
			zap.String("path", url.String()),
			zap.Any("params", req.Body),
		)

		resp, err := next(req)
		// 记录请求完成
		fields := []zap.Field{
			zap.String("requestID", reqId),
			zap.String("path", req.Raw.URL.Path),
			zap.Duration("latency", time.Since(startTime)),
		}

		if err != nil {
			logx.Logger.Error("Request failed", append(fields, zap.Error(err))...)
		} else {
			logx.Logger.Info("Request completed", fields...)
		}

		return resp, err
	}
}
