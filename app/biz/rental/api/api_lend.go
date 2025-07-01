package api

import (
	"context"
	"lend/gen/oas"
)

type LendHandler struct {
}

func NewLendHandler() *LendHandler {
	return &LendHandler{}
}

func (LendHandler) LendOut(ctx context.Context, request oas.OptLendOutReq) (*oas.LendOutOK, error) {
	return &oas.LendOutOK{}, nil

}
