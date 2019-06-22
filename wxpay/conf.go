package wxpay

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/henrylee2cn/goutil"
	"net/url"
	"strings"
)

type Config struct {
	AuthInfo        []*AuthInfo
}

type AuthInfo struct {
	Method    string
	AppID     string
	MchID     string
	Key       string
	AppSecret string
}

func (info *AuthInfo) MD5(data string) string {
	return strings.ToUpper(goutil.Md5([]byte(data + "&key=" + info.Key)))
}

func (info *AuthInfo) MD5V(data interface{}) string {
	v, _ := query.Values(data)
	d, _ := url.QueryUnescape(v.Encode())
	return info.MD5(d)
}

func (c *Config) FindAuthInfo(method string) (*AuthInfo, error) {
	for _, info := range c.AuthInfo {
		if info.Method == method {
			return info, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
