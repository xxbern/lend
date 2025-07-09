package api

import (
	"context"
	"lend/app/biz/user/internal/dba"
	"lend/gen/oas"
)

type UserManageHandler struct {
	userRepo *dba.UserRepository
}

func NewUserManageHandler(userRepo *dba.UserRepository) *UserManageHandler {
	return &UserManageHandler{userRepo: userRepo}
}

func (umh UserManageHandler) UserList(ctx context.Context, req oas.OptUserListReq) (oas.UserListRes, error) {
	list, err := umh.userRepo.UserList(ctx)

	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return &oas.UserListOK{
			Message: oas.NewOptString("Success"),
			Code:    oas.CommonCodeSUCCESS,
		}, nil
	}

	resUser := make([]oas.User, len(list))
	for i, item := range list {
		apiUser := &oas.User{
			ID:   oas.NewOptInt32(item.ID),
			Name: oas.NewOptString(*item.Name),
		}
		resUser[i] = *apiUser
	}

	return &oas.UserListOK{
		Message:  oas.NewOptString("Success"),
		Code:     oas.CommonCodeSUCCESS,
		UserList: resUser,
	}, nil
}
