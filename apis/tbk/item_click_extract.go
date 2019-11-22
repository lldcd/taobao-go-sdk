package tbk

import (
	"encoding/json"
	"fmt"
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

type ItemClickExtractResponse struct {
	ItemId string `json:"item_id"`
	OpenIid string `json:"open_iid"`
}

func (o *ItemClickExtract) StructToJson() string {
	json,err := json.Marshal(o)
	if err != nil{
		fmt.Println(err)
	}
	return string(json)
}

