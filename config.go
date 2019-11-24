package taobao_go_sdk

import (
	"github.com/nutsdo/taobao-go-sdk/utils"
	"net/http"
	"time"
)

const OPENURL = "https://eco.taobao.com/router/rest"

type Config struct {
	AutoRetry         bool            `default:"true"`
	MaxRetryTime      int             `default:"3"`
	UserAgent         string          `default:""`
	Debug             bool            `default:"false"`
	Timeout           time.Duration   `default:"10000000000"`
	HttpTransport     *http.Transport `default:""`
	EnableAsync       bool            `default:"false"`
	MaxTaskQueueSize  int             `default:"1000"`
	GoRoutinePoolSize int             `default:"5"`
	Scheme            string          `default:"HTTP"`
}

func NewConfig() (config *Config) {
	config = &Config{}
	utils.InitStructWithDefaultTag(config)
	return
}

func (c *Config) WithAutoRetry(isAutoRetry bool) *Config {
	c.AutoRetry = isAutoRetry
	return c
}

func (c *Config) WithMaxRetryTime(maxRetryTime int) *Config {
	c.MaxRetryTime = maxRetryTime
	return c
}

func (c *Config) WithUserAgent(userAgent string) *Config {
	c.UserAgent = userAgent
	return c
}

func (c *Config) WithDebug(isDebug bool) *Config {
	c.Debug = isDebug
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

func (c *Config) WithHttpTransport(httpTransport *http.Transport) *Config {
	c.HttpTransport = httpTransport
	return c
}

func (c *Config) WithEnableAsync(isEnableAsync bool) *Config {
	c.EnableAsync = isEnableAsync
	return c
}

func (c *Config) WithMaxTaskQueueSize(maxTaskQueueSize int) *Config {
	c.MaxTaskQueueSize = maxTaskQueueSize
	return c
}

func (c *Config) WithGoRoutinePoolSize(goRoutinePoolSize int) *Config {
	c.GoRoutinePoolSize = goRoutinePoolSize
	return c
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}
