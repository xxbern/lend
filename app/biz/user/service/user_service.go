package service

import (
	"context"
	"errors"
	"lend/gen/model"
)

var (
	ErrUserNotExist  = errors.New("user_not_exist")
	ErrInvalidCreds  = errors.New("invalid_credentials")
	ErrAccountLocked = errors.New("account_locked")
	// 可以轻松添加更多错误类型
)

// UserService 提供外部服务，除了该包，内部所有子包不允许被外部依赖
type (
	UserService interface {
		FindUser(ctx context.Context, id int32) (*model.UserInfo, error)

		DisableUser(ctx context.Context, id int32) error
	}
)
