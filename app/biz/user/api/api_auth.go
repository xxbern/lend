package api

import (
	"context"
	"github.com/ogen-go/ogen/conv"
	"go.uber.org/zap"
	"lend/app/auth"
	"lend/app/biz/user/dba"
	"lend/app/logx"
	"lend/gen/client/wx"
	"lend/gen/oas"
)

type WxLoginHandler struct {
	userRepo *dba.UserRepository
	wxClient *wx.Client
}

func NewWxLoginHandler(userRepo *dba.UserRepository, wxClient *wx.Client) *WxLoginHandler {
	return &WxLoginHandler{userRepo: userRepo, wxClient: wxClient}
}

func (sh WxLoginHandler) WxLogin(ctx context.Context, req oas.OptWxLoginReq) (oas.WxLoginRes, error) {
	wxLoginReq := req.Value
	uid, _ := conv.ToInt32(wxLoginReq.LoginCode)
	user, _ := sh.userRepo.FindById(ctx, uid)
	logx.Logger.Info("user info", zap.Any("user", user))

	tokenUser := auth.TokenUser{UserInfo: user}
	tokenStr, _ := tokenUser.NewToken()

	apiUser := oas.NewOptUser(oas.User{
		ID: oas.NewOptInt32(user.ID),
	})
	response, err := sh.wxClient.Code2Session(ctx, wx.Code2SessionParams{JsCode: "asdf", Appid: "asdf"})
	if err != nil {
		logx.Logger.Error("code2Session", zap.Error(err))
		return nil, err
	}
	logx.Logger.Error("code2Session", zap.Any("response", response))

	return &oas.WxLoginOK{
		Message:     oas.NewOptString("Login Success"),
		Code:        oas.OptWxLoginOKCode{Value: oas.WxLoginOKCodeSUCCESS, Set: true},
		AccessToken: oas.NewOptString(tokenStr),
		UserInfo:    apiUser,
	}, nil
}

func (WxLoginHandler) WxPhone(ctx context.Context, req oas.OptWxPhoneReq) (oas.WxPhoneRes, error) {
	//TODO implement me
	panic("implement me")
}
