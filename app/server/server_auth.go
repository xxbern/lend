package server

import (
	"context"
	set "github.com/deckarep/golang-set/v2"
	"github.com/ogen-go/ogen/ogenerrors"
	"lend/app/auth"
	"lend/gen/oas"
)

//func auth() middleware.Middleware {
//	return func(
//		req middleware.Request,
//		next func(req middleware.Request) (middleware.Response, error),
//	) (middleware.Response, error) {
//
//		//err := reflect.ValueOf(req.Body).MethodByName("Validate").Call(nil)
//		//
//		//if err != nil {
//		//	log.Fatal("Fail", err)
//		//}
//
//		constant := context.WithValue(req.Context, CtxKeyUserID, "i'm a userId")
//		req.SetContext(constant)
//		token := req.Raw.Header.Get("Authorization")
//
//		if token == "" {
//			// 直接拦截请求，返回 401 Unauthorized 错误响应
//			return middleware.Response{
//				Type: &oas.CommonResponseStatusCode{StatusCode: 401, Response: oas.CommonResponse{Code: oas.CommonCodeUNAUTHORIZED, Message: oas.NewOptString("登录失效")}},
//			}, nil
//		}
//
//		resp, e := next(req)
//
//		return resp, e
//	}
//
//}

type SecurityHandler struct{}

func (SecurityHandler) HandleBearer(ctx context.Context, operationName oas.OperationName, t oas.Bearer) (context.Context, error) {
	if t.Token == "" {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}
	user, roles, err := auth.Decode2User(t.Token)

	if err != nil {
		return ctx, ogenerrors.ErrSkipServerSecurity
	}
	withUser := auth.WithAuthedUser(ctx, &auth.TokenUser{UserInfo: user})
	if len(t.Roles) == 0 {
		return withUser, nil
	}
	userRoles := set.NewSet(roles...)
	if !userRoles.Contains(t.Roles...) {
		return ctx, ogenerrors.ErrSecurityRequirementIsNotSatisfied
	}
	return withUser, nil
}
