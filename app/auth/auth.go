package auth

import (
	"context"
	"lend/app/config"
)

type key string

const (
	KeyAuthedUser key = "authedUser"
)

func WithAuthedUser(ctx context.Context, user *TokenUser) context.Context {
	return context.WithValue(ctx, KeyAuthedUser, user)
}

func AuthedUser(ctx context.Context) *TokenUser {
	return (ctx.Value(config.KeyAuthedUser)).(*TokenUser)
}
