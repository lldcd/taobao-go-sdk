package main

import (
	"fmt"
	"github.com/nutsdo/taobao-go-sdk/apis/tbk"
)

func main() {

	//app := auth.TaoBaoApp{AppKey:"25254678",AppSecret:"5e5e16ddb6c2d780f91401a9c990b7d6"}
	//api := &tbk.ItemClickExtract{}
	//api.SetClickUrl("https://s.click.taobao.com/t?e=m%3D2%26s%3D3WASVFHrRukcQipKwQzePOeEDrYVVa64K7Vc7tFgwiHjf2vlNIV67rEkNF%2FKS%2FS5myBzYSO0LNzT6y8UW0x6Ul%2FZMKpynZSgUyMoWRC37ckebvs5ByqxLPAy%2Fay3dFHhPplz8soEI3BvDHPhpDocWpNhaOS5%2Fa2OOemaFM5tHHYxZyjQcbVDhcnjRDTsxzJ61dvbXYNAZVeVpXSFzjVS3MYOae24fhW0&scm=null&pvid=null&app_pvid=59590_11.1.93.222_519_1574413600225&ptl=floorId:17741;originalFloorId:17741;app_pvid:59590_11.1.93.222_519_1574413600225&union_lens=lensId%3AOPT%401574413590%400b1a25ec_0e4e_16e925ce072_5034%4001")
	//请求
	//resData, respErr := client.DefaultClient.DoRequest("POST","", app.BuildQuery(api))
	//if respErr != nil {
	//	fmt.Println(respErr)
	//}
	//fmt.Println(string(resData))
	//var ItemClickExtractResponse  = tbk.ItemClickExtractResponse{}
	//json.Unmarshal(resData,&ItemClickExtractResponse)
	//
	//fmt.Printf("%#v", ItemClickExtractResponse)

	SdkClient, err := tbk.NewClientWithAccessKey("25254678", "5e5e16ddb6c2d780f91401a9c990b7d6")
	if err != nil {
		fmt.Println("create client err:", err.Error())
	}
	req := tbk.TaobaoTbkItemClickExtractRequest("https://s.click.taobao.com/t?e=m%3D2%26s%3D3WASVFHrRukcQipKwQzePOeEDrYVVa64K7Vc7tFgwiHjf2vlNIV67rEkNF%2FKS%2FS5myBzYSO0LNzT6y8UW0x6Ul%2FZMKpynZSgUyMoWRC37ckebvs5ByqxLPAy%2Fay3dFHhPplz8soEI3BvDHPhpDocWpNhaOS5%2Fa2OOemaFM5tHHYxZyjQcbVDhcnjRDTsxzJ61dvbXYNAZVeVpXSFzjVS3MYOae24fhW0&scm=null&pvid=null&app_pvid=59590_11.1.93.222_519_1574413600225&ptl=floorId:17741;originalFloorId:17741;app_pvid:59590_11.1.93.222_519_1574413600225&union_lens=lensId%3AOPT%401574413590%400b1a25ec_0e4e_16e925ce072_5034%4001")

	res, err := SdkClient.TaobaoTbkItemClickExtract(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res:", res.ItemId)
}
