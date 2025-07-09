package internal

import (
	"context"
	"lend/app/biz/user/service"
	"lend/gen/model"
)

type UserServiceInternal interface {
	service.UserService

	UserWxLogin(ctx context.Context, wxAppid string, wxAuthCode string) (tokenStr string, user *model.UserInfo, err error)
}
