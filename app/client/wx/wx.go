package wx

import "lend/gen/client/wx"

func Init() (*wx.Client, error) {
	return wx.NewClient("https://api.weixin.qq.com")
}
