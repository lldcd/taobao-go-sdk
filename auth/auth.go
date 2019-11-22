package auth

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
	"github.com/nutsdo/taobao-go-sdk/apis"
	"io"
	"net/url"
	"sort"
	"strings"
	"time"
)

type TaoBaoApp struct {
	AppKey string
	AppSecret string
}

func New(appkey,secret string) *TaoBaoApp {
	return &TaoBaoApp{AppKey:appkey,AppSecret:secret}
}

func (app *TaoBaoApp) Signature(params map[string]interface{}) string {

	var keys []string
	params["app_key"] = app.AppKey
	for k, v := range params{
		switch v.(type) {
		case []byte:
		default:
			keys = append(keys, k)
		}
	}
	//数组按ASCII字典升序排序
	sort.Strings(keys)
	//拼接字符串
	var query strings.Builder

	if params["sign_method"] == "md5" {

		query.WriteString(app.AppSecret)
	}
	for _, key := range keys {
		query.WriteString(key)
		v,_:=params[key].(string)
		if v!="" {
			query.WriteString(v)
		}

	}
	fmt.Println(query.String())
	var sign string
	if params["sign_method"] == "hmac" {
		//生成签名方式:HMAC_MD5
		sign = fmt.Sprintf("%x", hmac.New(md5.New,[]byte(app.AppSecret)))
	}else {
		//生成签名方式:MD5
		query.WriteString(app.AppSecret)
		sign = fmt.Sprintf("%x", md5.Sum([]byte(query.String())))
	}
	fmt.Println(sign)
	return strings.ToUpper(sign)
}

func (app *TaoBaoApp) BuildQuery(api apis.TaobaoApiInterface) io.Reader {

	paramsMap := make(map[string]interface{})
	paramsMap["method"] = api.GetMethod()
	paramsMap["app_key"]= app.AppKey
	paramsMap["sign_method"] = "md5"
	paramsMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	paramsMap["format"] = "json"
	paramsMap["v"] = "2.0"
	for k,_:=range api.GetValues(){
		paramsMap[k] = api.GetValues().Get(k)
	}
	query := url.Values{}
	for k, v := range paramsMap {
		query.Add(k,v.(string))
	}

	query.Add("sign",app.Signature(paramsMap))

	return strings.NewReader(query.Encode())
}

