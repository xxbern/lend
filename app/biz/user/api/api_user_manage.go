package api

import (
	"context"
	"lend/gen/oas"
)

type UserManageHandler struct {
}

func NewUserManageHandler() *UserManageHandler {
	return &UserManageHandler{}
}

func (UserManageHandler) UserList(ctx context.Context, req oas.OptUserListReq) (oas.UserListRes, error) {
	return &oas.UserListOK{
		Message: oas.NewOptString("Success"),
		Code:    oas.CommonCodeSUCCESS,
	}, nil
}
