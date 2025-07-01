package api

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ogen-go/ogen/ogenerrors"
	"lend/app/auth"
	"lend/app/biz/user/internal/dba"
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
	apiUser := oas.User{}

	err := copier.Copy(&apiUser, &user)
	if err != nil {
		return nil, err
	}

	return &oas.UserInfoOK{
		Message:  oas.NewOptString("Success"),
		Code:     oas.CommonCodeSUCCESS,
		UserInfo: oas.NewOptUser(apiUser),
	}, nil
}
