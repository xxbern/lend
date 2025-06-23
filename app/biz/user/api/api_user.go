package api

import (
	"context"
	"github.com/ogen-go/ogen/ogenerrors"
	"lend/app/auth"
	"lend/app/biz/user/dba"
	"lend/gen/client/wx"
	"lend/gen/oas"
)

type UserInfoHandler struct {
	userRepo *dba.UserRepository
	wxClient *wx.Client
}

func NewUserInfoHandler(wxClient *wx.Client, userRepo *dba.UserRepository) *UserInfoHandler {
	return &UserInfoHandler{wxClient: wxClient, userRepo: userRepo}
}

func (sh UserInfoHandler) UserInfo(ctx context.Context) (oas.UserInfoRes, error) {
	authedUser := auth.AuthedUser(ctx)
	if authedUser == nil {
		return nil, ogenerrors.ErrSkipServerSecurity
	}
	user, e := sh.userRepo.FindById(ctx, authedUser.ID)
	if e != nil {
		return nil, e
	}

	apiUser := oas.NewOptUser(oas.User{
		ID: oas.NewOptInt32(user.ID),
	})

	return &oas.UserInfoOK{
		Message:  oas.NewOptString("Success"),
		Code:     oas.CommonCodeSUCCESS,
		UserInfo: apiUser,
	}, nil
}
