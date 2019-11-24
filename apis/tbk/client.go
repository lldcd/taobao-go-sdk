package tbk

import (
	taobao_go_sdk "github.com/nutsdo/taobao-go-sdk"
)

type Client struct {
	taobao_go_sdk.Client
}

func NewClientWithAccessKey(accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(accessKeyId, accessKeySecret)
	return
}
