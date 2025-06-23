package main

import "testing"

func Test(t *testing.T) {
	GenGorm()
	GenApi()
}

//go:generate ogen -config gen-api-wx-config.yml --target ../gen/client/wx --package wx -clean ../api/client/openapi-wx.yml
