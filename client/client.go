package client

import (
	"io"
	"io/ioutil"
	"net/http"
)

var (
	DefaultClient = Client{&http.Client{Transport:http.DefaultTransport}}
	OpenUrl = "https://eco.taobao.com/router/rest"
	)


type Client struct {
	*http.Client
}

func (c Client) DoRequest(method,reqUrl string,body io.Reader) (resData []byte,err error) {

	if reqUrl == "" {
		reqUrl = OpenUrl
	}
	req,err := http.NewRequest("POST",reqUrl, body)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := c.Do(req)
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

