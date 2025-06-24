package auth

import (
	"context"
)

type key string

const (
	KeyAuthedUser key = "authedUser"
)

func WithAuthedUser(ctx context.Context, user *TokenUser) context.Context {
	return context.WithValue(ctx, KeyAuthedUser, user)
}

func AuthedUser(ctx context.Context) *TokenUser {
	return (ctx.Value(KeyAuthedUser)).(*TokenUser)
}
