package wxpay

import (
	"encoding/xml"
	"fmt"
	"github.com/hzxiao/goutil"
	"github.com/hzxiao/goutil/httputil"
	"github.com/hzxiao/goutil/log"
	"io/ioutil"
	"net/http"
)

var AccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
var JsCode2SessionUrl = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
var UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder"
var QueryOrderUrl = "https://api.mch.weixin.qq.com/pay/orderquery"

const (
	TradeTypeJSAPI  = "JSAPI"
	TradeTypeNative = "NATIVE"
	TradeTypeAPP    = "APP"
	TradeTypeMWeb   = "MWEB"
)

type NotifyHandler interface {
	PayMethod() string
	OnNotify(info *AuthInfo, req *NotifyRequest) error
}

// Client WeChat pay client
type Client struct {
	config        *Config
	notifyHandler NotifyHandler
}

func NewClient(config *Config, handler NotifyHandler) *Client {
	return &Client{
		config:        config,
		notifyHandler: handler,
	}
}

// LoadOpenID load open id by code
func (c *Client) LoadOpenID(payMethod, code string, mp bool) (goutil.Map, error) {
	authInfo, err := c.config.FindAuthInfo(payMethod)
	if err != nil {
		return nil, fmt.Errorf("find auth info by method(%v) err: %v", payMethod, err)
	}

	var url = AccessTokenUrl
	if mp {
		url = JsCode2SessionUrl
	}
	var result goutil.Map
	err = httputil.Get(fmt.Sprintf(url, authInfo.AppID, authInfo.AppSecret, code), httputil.ReturnJSON, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UnifiedOrder(method string, req *OrderUnifiedRequest) (*OrderUnifiedResponse, error) {
	info, err := c.config.FindAuthInfo(method)
	if err != nil {
		return nil, fmt.Errorf("wxpay: UnifiedOrder find auth info by %v err: %v", method, err)
	}

	req.AppID = info.AppID
	req.MchID = info.Method
	req.SetSign(info)

	var response OrderUnifiedResponse
	err = httputil.PostXML(UnifiedOrderUrl, req, httputil.ReturnXML, &response)
	if err != nil {
		return nil, err
	}

	if response.ReturnCode != "SUCCESS" {
		return nil, fmt.Errorf("wxpay: UnifiedOrder by data(\n%v\n) fail with code(%v)error(%v)->%v", req, response.ReturnCode, response.ReturnMsg, response)
	}

	if !response.VerifySign(info, response.Sign) {
		return nil, fmt.Errorf("wxpay: UnifiedOrder verify sign with data(\n%v\n) fail", response)
	}
	return &response, nil
}

func (c *Client) QueryOrder(method string, req *OrderQueryRequest) (*OrderQueryResponse, error) {
	info, err := c.config.FindAuthInfo(method)
	if err != nil {
		return nil, fmt.Errorf("wxpay: QueryOrder find auth info by %v err: %v", method, err)
	}

	req.AppID = info.AppID
	req.MchID = info.Method
	req.SetSign(info)

	var response OrderQueryResponse
	err = httputil.PostXML(QueryOrderUrl, req, httputil.ReturnXML, &response)
	if err != nil {
		return nil, err
	}

	if response.ReturnCode != "SUCCESS" {
		return nil, fmt.Errorf("wxpay: QueryOrder by data(\n%v\n) fail with code(%v)error(%v)->%v", req, response.ReturnCode, response.ReturnMsg, response)
	}

	if !response.VerifySign(info, response.Sign) {
		return nil, fmt.Errorf("wxpay: QueryOrder verify sign with data(\n%v\n) fail", response)
	}
	return &response, nil
}

func (c *Client) OnNotify(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp = &NotifyResponse{}
	defer func() {
		if err != nil {
			resp.ReturnCode = "FAIL"
			resp.ReturnMsg = err.Error()
		} else {
			resp.ReturnCode = "SUCCESS"
			resp.ReturnMsg = "OK"
		}
		bys, _ := xml.Marshal(resp)
		w.Write(bys)
	}()

	info, err := c.config.FindAuthInfo(c.notifyHandler.PayMethod())
	if err != nil {
		err = fmt.Errorf("wxpay: OnNotify find auth info by %v err: %v", c.notifyHandler.PayMethod(), err)
		log.Error("%v", err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("wxpay: OnNotify read req body err: %v", err)
		log.Error("%v", err)
		return
	}
	defer r.Body.Close()

	var notification NotifyRequest
	err = xml.Unmarshal(body, &notification)
	if err != nil {
		err = fmt.Errorf("wxpay: OnNotify decode XML body(\n%v\n) err: %v", body, err)
		log.Error("%v", err)
		return
	}

	if !notification.VerifySign(info, notification.Sign) {
		err = fmt.Errorf("wxpay: OnNotify verify sign body(\n%v\n) err: verify sign fail", body)
		log.Error("%v", err)
		return
	}

	err = c.notifyHandler.OnNotify(info, &notification)
	if err != nil {
		err = fmt.Errorf("wxpay: OnNotify handle notify body(\n%v\n) err: %v", body, err)
		log.Error("%v", err)
		return
	}
}
