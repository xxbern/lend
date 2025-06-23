package server

import (
	"context"
	"errors"
	"github.com/ogen-go/ogen/ogenerrors"
	"lend/gen/oas"
	"net/http"
)

func errorHandle() ogenerrors.ErrorHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {

		var secErr *ogenerrors.SecurityError
		var resp oas.CommonResponse
		if errors.As(err, &secErr) {
			resp = oas.CommonResponse{Code: oas.CommonCodeUNAUTHORIZED, Message: oas.NewOptString("未登录")}
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			resp = oas.CommonResponse{Code: oas.CommonCodeUNKNOWN, Message: oas.NewOptString("未知异常")}
			w.WriteHeader(http.StatusInternalServerError)
		}

		json, _ := resp.MarshalJSON()
		_, _ = w.Write(json)
	}
}
