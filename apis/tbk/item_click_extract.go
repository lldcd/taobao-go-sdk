package tbk

import (
	"encoding/json"
	"fmt"
	"github.com/nutsdo/taobao-go-sdk/apis"
	"github.com/nutsdo/taobao-go-sdk/responses"
)

type ItemClickExtract struct {
	apis.TaobaoApi
}

func (t *ItemClickExtract) GetMethod() string {
	return "taobao.tbk.item.click.extract"
}

func (t *ItemClickExtract) SetClickUrl(clickUrl string) {
	//fmt.Println("item_click_extract->SetClickUrl->ItemClickExtract:",t)
	t.SetValue("click_url", clickUrl)
}

type ItemClickExtractResponse struct {
	ItemId  string `json:"item_id"`
	OpenIid string `json:"open_iid"`
}

type ItemClickExtractFullResponse struct {
	*responses.BaseResponse
	ItemId  string `json:"item_id"`
	OpenIid string `json:"open_iid"`
}

func (o *ItemClickExtract) StructToJson() string {
	json, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	return string(json)
}

//sdk测试
func (client *Client) TaobaoTbkItemClickExtract(request *ItemClickExtract) (response *ItemClickExtractFullResponse, err error) {
	response = CreateTaobaoTbkItemClickExtractResonse()
	resData, err := client.DoRequestFull("POST", "", client.BuildQuery(request))
	if err != nil {
		return
	}
	jsonErr := responses.Unmarshal(response, resData, "json")
	if jsonErr != nil {
		err = jsonErr
	}
	return
}

func TaobaoTbkItemClickExtractRequest(url string) (request *ItemClickExtract) {
	request = &ItemClickExtract{}
	request.SetClickUrl(url)
	return
}

func CreateTaobaoTbkItemClickExtractResonse() (response *ItemClickExtractFullResponse) {
	response = &ItemClickExtractFullResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
