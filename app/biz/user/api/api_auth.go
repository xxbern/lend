package api

import (
	"context"
	"errors"
	"lend/app/biz/user/internal"
	"lend/app/biz/user/internal/dba"
	"lend/app/biz/user/service"
	"lend/gen/client/wx"
	"lend/gen/oas"
)

type WxLoginHandler struct {
	userRepo    *dba.UserRepository
	wxClient    *wx.Client
	userService internal.UserServiceInternal
}

func NewWxLoginHandler(userService internal.UserServiceInternal, wxClient *wx.Client, userRepo *dba.UserRepository) *WxLoginHandler {
	return &WxLoginHandler{userService: userService, wxClient: wxClient, userRepo: userRepo}
}

func (sh WxLoginHandler) WxLogin(ctx context.Context, wxLoginReq *oas.WxLoginReq) (oas.WxLoginRes, error) {
	token, userInfo, err := sh.userService.UserWxLogin(ctx, wxLoginReq.Appid, wxLoginReq.LoginCode)
	if errors.As(err, &service.ErrUserNotExist) {
		return &oas.WxLoginOK{
			Message: oas.NewOptString("当前微信号未关联账户"),
			Code:    oas.OptWxLoginOKCode{Value: oas.WxLoginOKCodeUNBINDPHONE, Set: true},
		}, nil
	}
	if err != nil {
		return nil, err
	}

	apiUser := oas.NewOptUser(oas.User{
		ID: oas.NewOptInt32(userInfo.ID),
	})

	return &oas.WxLoginOK{
		Message:     oas.NewOptString("Login Success"),
		Code:        oas.OptWxLoginOKCode{Value: oas.WxLoginOKCodeSUCCESS, Set: true},
		AccessToken: oas.NewOptString(token),
		UserInfo:    apiUser,
	}, nil
}

func (WxLoginHandler) WxPhone(ctx context.Context, req *oas.WxPhoneReq) (oas.WxPhoneRes, error) {
	//TODO implement me
	panic("implement me")
}
