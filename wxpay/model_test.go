package wxpay

import (
	"encoding/xml"
	"github.com/hzxiao/goutil/assert"
	"github.com/k0kubun/pp"
	"testing"
)

var auth *AuthInfo

func TestMain(m *testing.M) {
	auth = &AuthInfo{
		Method: "mp",
		AppID:  "wx2421b1c4370ec43b",
		MchID:  "10000100",
		Key:    "123456",
	}
	m.Run()
}

func TestOrderUnifiedRequest_SetSign(t *testing.T) {

	s := `<xml>
   <appid>wx2421b1c4370ec43b</appid>
   <attach>支付测试</attach>
   <body>JSAPI支付测试</body>
   <mch_id>10000100</mch_id>
   <detail><![CDATA[{ "goods_detail":[ { "goods_id":"iphone6s_16G", "wxpay_goods_id":"1001", "goods_name":"iPhone6s 16G", "quantity":1, "price":528800, "goods_category":"123456", "body":"苹果手机" }, { "goods_id":"iphone6s_32G", "wxpay_goods_id":"1002", "goods_name":"iPhone6s 32G", "quantity":1, "price":608800, "goods_category":"123789", "body":"苹果手机" } ] }]]></detail>
   <nonce_str>1add1a30ac87aa2db72f57a2375d8fec</nonce_str>
   <notify_url>http://wxpay.wxutil.com/pub_v2/pay/notify.v2.php</notify_url>
   <openid>oUpF8uMuAJO_M2pxb1Q9zNjWeS6o</openid>
   <out_trade_no>1415659990</out_trade_no>
   <spbill_create_ip>14.23.150.211</spbill_create_ip>
   <total_fee>1</total_fee>
   <trade_type>JSAPI</trade_type>
   <sign>0CB01533B8C1EF103065174F50BCA001</sign>
</xml>`
	r := &OrderUnifiedRequest{}
	err := xml.Unmarshal([]byte(s), r)
	assert.NoError(t, err)

	r.SetSign(auth)

	assert.NotEqual(t, "", r.Sign)
}

func TestOrderUnifiedResponse_VerifySign(t *testing.T) {
	s := `<xml>
   <return_code><![CDATA[SUCCESS]]></return_code>
   <return_msg><![CDATA[OK]]></return_msg>
   <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
   <mch_id><![CDATA[10000100]]></mch_id>
   <nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str>
   <openid><![CDATA[oUpF8uMuAJO_M2pxb1Q9zNjWeS6o]]></openid>
   <sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign>
   <result_code><![CDATA[SUCCESS]]></result_code>
   <prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id>
   <trade_type><![CDATA[JSAPI]]></trade_type>
</xml>`

	r := &OrderUnifiedResponse{}
	err := xml.Unmarshal([]byte(s), r)
	assert.NoError(t, err)

	pp.Println(r)
	sign := auth.MD5V(r)
	assert.True(t, r.VerifySign(auth, sign))
}