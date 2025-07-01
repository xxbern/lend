package internal

import (
	"context"
	"go.uber.org/zap"
	"lend/app/auth"
	"lend/app/biz/user/internal/dba"
	"lend/app/biz/user/service"
	"lend/app/logx"
	"lend/gen/client/wx"
	"lend/gen/model"
)

// 如果在业务根目录提供了接口，在此提供相应的实现

type UserServiceImpl struct {
	userRepo *dba.UserRepository
	wxClient *wx.Client
}

func NewUserServiceImpl(wxClient *wx.Client, userRepo *dba.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{wxClient: wxClient, userRepo: userRepo}
}

func (u UserServiceImpl) DisableUser(ctx context.Context, id int32) error {
	return u.userRepo.DisableUser(ctx, id)
}

func (u UserServiceImpl) UserWxLogin(ctx context.Context, wxAppid string, wxAuthCode string) (tokenStr string, user *model.UserInfo, err error) {
	session, err := u.wxClient.Code2Session(
		ctx,
		wx.Code2SessionParams{
			Appid:     wxAppid,
			Secret:    "",
			JsCode:    wxAuthCode,
			GrantType: wx.Code2SessionGrantTypeAuthorizationCode,
		},
	)
	logx.Logger.Info("wxLogin", zap.Any("resp", session), zap.Error(err))
	if err != nil {
		return "", nil, err
	}
	optOpenId := session.GetOpenid()

	uf, err := u.userRepo.FindByForeignerId(ctx, wxAppid, optOpenId.Value)
	if err != nil {
		return "", nil, err
	}
	if uf == nil || uf.UserID == 0 {
		return "", nil, service.ErrUserNotExist
	}

	tokenUser := auth.TokenUser{UserInfo: &uf.UserInfo}
	tokenStr, err = tokenUser.NewToken()
	if err != nil {
		return "", nil, err
	}

	return tokenStr, &uf.UserInfo, nil
}

func (u UserServiceImpl) FindUser(ctx context.Context, id int32) (*model.UserInfo, error) {
	return u.userRepo.FindById(ctx, id)
}

func (u UserServiceImpl) Asdf(ctx context.Context, id int32) (*model.UserInfo, error) {
	return u.userRepo.FindById(ctx, id)
}
