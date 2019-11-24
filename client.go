package taobao_go_sdk

import (
	"github.com/nutsdo/taobao-go-sdk/apis"
	"github.com/nutsdo/taobao-go-sdk/auth"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	//DefaultClient = Client{&http.Client{Transport:http.DefaultTransport}}
	OpenUrl = "https://eco.taobao.com/router/rest"
)

type Client struct {
	//*http.Client
	httpClient *http.Client
	isRunning  bool
	config     *Config
	credential *auth.TaoBaoApp
}

func (c Client) DoRequest(method, reqUrl string, body io.Reader) (resData []byte, err error) {

	if reqUrl == "" {
		reqUrl = OpenUrl
	}
	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := c.httpClient.Do(req)
	if respErr != nil {
		err = respErr
		return
	}
	defer resp.Body.Close()

	resData, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		err = ioErr
		return
	}
	return
}

func (c Client) DoRequestFull(method, reqUrl string, body io.Reader) (resp *http.Response, err error) {

	if reqUrl == "" {
		reqUrl = OpenUrl
	}
	req, err := http.NewRequest("POST", reqUrl, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := c.httpClient.Do(req)
	if respErr != nil {
		err = respErr
		return
	}
	return
}

func (client *Client) BuildQuery(api apis.TaobaoApiInterface) io.Reader {

	paramsMap := make(map[string]interface{})
	paramsMap["method"] = api.GetMethod()
	paramsMap["app_key"] = client.credential.AppKey
	paramsMap["sign_method"] = "md5"
	paramsMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	paramsMap["format"] = "json"
	paramsMap["v"] = "2.0"
	if paramsMap["format"] == "json" {
		paramsMap["simplify"] = "true"
	}
	for k, _ := range api.GetValues() {
		paramsMap[k] = api.GetValues().Get(k)
	}
	query := url.Values{}
	for k, v := range paramsMap {
		switch v.(type) {
		case string:
			query.Add(k, v.(string))
		default:
		}
		if v, ok := v.(string); ok {
			query.Add(k, v)
		}
	}

	query.Add("sign", client.credential.Signature(paramsMap))

	return strings.NewReader(query.Encode())
}

func NewClientWithAccessKey(accessKeyId, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(accessKeyId, accessKeySecret)
	return
}

func (client *Client) InitWithAccessKey(accessKeyId, accessKeySecret string) (err error) {
	config := client.InitClientConfig()
	credential := &auth.TaoBaoApp{
		AppKey:    accessKeyId,
		AppSecret: accessKeySecret,
	}
	//fmt.Println("client->InitWithAccessKey:",*credential)
	err = client.InitWithOptions(config, credential)

	return
}

func (client *Client) InitWithOptions(config *Config, credential *auth.TaoBaoApp) (err error) {
	client.isRunning = true
	client.httpClient = &http.Client{}
	client.credential = credential
	//fmt.Println("client->InitWithOptions->credential:",credential.AppKey)

	if config.HttpTransport != nil {
		client.httpClient.Transport = config.HttpTransport
	}

	if config.Timeout > 0 {
		client.httpClient.Timeout = config.Timeout
	}

	//client.signer, err = auth.NewSignerWithCredential(credential, client.ProcessCommonRequestWithSigner)

	return
}

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	} else {
		return NewConfig()
	}
}
