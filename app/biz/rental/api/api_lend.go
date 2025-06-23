package api

import (
	"context"
	"lend/app/biz/user/dba"
	"lend/gen/oas"
)

type LendHandler struct {
	userRepo *dba.UserRepository
}

func NewLendHandler() *LendHandler {
	return &LendHandler{}
}

func (LendHandler) LendOut(ctx context.Context, request oas.OptLendOutReq) (*oas.LendOutOK, error) {

	return &oas.LendOutOK{}, nil

}
