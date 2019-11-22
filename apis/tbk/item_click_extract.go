package tbk

import (
	"github.com/nutsdo/taobao-go-sdk/apis"
)

type ItemClickExtract struct {
	apis.TaobaoApi
}

func (t *ItemClickExtract) GetMethod() string {
	return "taobao.tbk.item.click.extract"
}

func (t *ItemClickExtract) SetClickUrl(clickUrl string) {
	t.SetValue("click_url",clickUrl)
}