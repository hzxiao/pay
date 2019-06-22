package wxpay

import (
	"encoding/xml"
	"github.com/hzxiao/goutil"
	"strings"
)

type OrderUnifiedRequest struct {
	XMLName        xml.Name `xml:"xml" url:"-"`
	AppID          string   `xml:"appid,omitempty" url:"appid,omitempty"`                       // appid	 appid  是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
	MchID          string   `xml:"mch_id,omitempty" url:"mch_id,omitempty"`                     // 商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	NonceStr       string   `xml:"nonce_str,omitempty" url:"nonce_str,omitempty"`               // 随机字符串	nonce_str 	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
	Sign           string   `xml:"sign,omitempty" url:"-"`                                      // 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
	Body           string   `xml:"body,omitempty" url:"body,omitempty"`                         // 商品描述	body	是	String(32)	Ipad mini  16G  白色	商品或支付单简要描述
	OutTradeNo     string   `xml:"out_trade_no,omitempty" url:"out_trade_no,omitempty"`         // 商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	TotalFee       int      `xml:"total_fee,omitempty" url:"total_fee,omitempty"`               // 总金额	 total_fee	是	Int	888	订单总金额，单位为分，详见支付金额
	SpbillCreateIp string   `xml:"spbill_create_ip,omitempty" url:"spbill_create_ip,omitempty"` // 终端IP	spbill_create_ip	是	String(16)	123.12.12.123	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	NotifyUrl      string   `xml:"notify_url,omitempty" url:"notify_url,omitempty"`             // 通知地址	notify_url	是	String(256)	http://www.weixin.qq.com/wxpay/pay.php	接收微信支付异步通知回调地址
	TradeType      string   `xml:"trade_type,omitempty" url:"trade_type,omitempty"`             // 交易类型	trade_type	是	String(16)	JSAPI	取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
	ProductId      string   `xml:"product_id,omitempty" url:"product_id,omitempty"`             // 商品ID	product_id	否	String(32)	12235413214070356458058	trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
	DeviceInfo     string   `xml:"device_info,omitempty" url:"device_info,omitempty"`           // 设备号	device_info	否	String(32)	013467007045764	终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	Detail         string   `xml:"detail,omitempty" url:"detail,omitempty"`                     // 商品详情	detail	否	String(8192)	Ipad mini  16G  白色	商品名称明细列表
	Attach         string   `xml:"attach,omitempty" url:"attach,omitempty"`                     // 附加数据	attach	否	String(127)	深圳分店	附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType        string   `xml:"fee_type,omitempty" url:"fee_type,omitempty"`                 // 货币类型	fee_type	否	String(16)	CNY	符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart      string   `xml:"time_start,omitempty" url:"time_start,omitempty"`             // 交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire     string   `xml:"time_expire,omitempty" url:"time_expire,omitempty"`           // 交易结束时间	time_expire	否	String(14)	20091227091010, 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则,注意：最短失效时间间隔必须大于5分钟
	GoodsTag       string   `xml:"goods_tag,omitempty" url:"goods_tag,omitempty"`               // 商品标记	goods_tag	否	String(32)	WXG	商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	LimitPay       string   `xml:"limit_pay,omitempty" url:"limit_pay,omitempty"`               // 指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
	Openid         string   `xml:"openid,omitempty" url:"openid,omitempty"`                     // 用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
}

//
//func (o *OrderArgs) SetDetail(goods []*Goods) {
//	o.Detail = util.S2Json(util.Map{
//		"goods_detail": goods,
//	})
//}
//
func (o *OrderUnifiedRequest) SetSign(info *AuthInfo) {
	o.NonceStr = strings.ToUpper(goutil.UUID())
	o.Sign = info.MD5V(o)
}

type OrderUnifiedResponse struct {
	ReturnCode string `xml:"return_code" url:"return_code,omitempty"`   // 返回状态码	return_code	是	String(16)	SUCCESS
	ReturnMsg  string `xml:"return_msg" url:"return_msg,omitempty"`     // 返回信息	return_msg	否	String(128)	签名失败
	Appid      string `xml:"appid" url:"appid,omitempty"`               // 公众账号ID	appid	是	String(32)	wx8888888888888888	调用接口提交的公众账号ID
	MchId      string `xml:"mch_id" url:"mch_id,omitempty"`             // 商户号	mch_id	是	String(32)	1900000109	调用接口提交的商户号
	DeviceInfo string `xml:"device_info" url:"device_info,omitempty"`   // 设备号	device_info	否	String(32)	013467007045764	调用接口提交的终端设备号，
	NonceStr   string `xml:"nonce_str" url:"nonce_str,omitempty"`       // 随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	微信返回的随机字符串
	Sign       string `xml:"sign" url:"-"`                              // 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	微信返回的签名，详见签名算法
	ResultCode string `xml:"result_code" url:"result_code,omitempty"`   // 业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
	ErrCode    string `xml:"err_code" url:"err_code,omitempty"`         // 错误代码	err_code	否	String(32)	SYSTEMERROR	详细参见第6节错误列表
	ErrCodeDes string `xml:"err_code_des" url:"err_code_des,omitempty"` // 错误代码描述	err_code_des	否	String(128)	系统错误	错误返回的信息描述
	PrepayId   string `xml:"prepay_id" url:"prepay_id,omitempty"`       // 预支付交易会话标识	prepay_id	是	String(64)	wx201410272009395522657a690389285100	微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	TradeType  string `xml:"trade_type" url:"trade_type,omitempty"`     // 交易类型	trade_type	是	String(16)	JSAPI	调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
	CodeUrl    string `xml:"code_url" url:"code_url,omitempty"`         // 二维码链接	code_url	否	String(64)	URl：weixin：//wxpay/s/An4baqw	trade_type为NATIVE是有返回，可将该参数值生成二维码展示出来进行扫码支付
}

func (o OrderUnifiedResponse) VerifySign(info *AuthInfo, sign string) bool {
	if info.MD5V(o) == sign {
		return true
	}
	return false
}

type OrderQueryRequest struct {
	// 	公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
	AppID string `xml:"appid" url:"appid,omitempty"`
	// 商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	MchID string `xml:"mch_id" url:"mch_id,omitempty"`
	// 微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，优先使用
	TransactionId string `xml:"transaction_id" url:"transaction_id,omitempty"`
	// 商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
	OutTradeNo string `xml:"out_trade_no" url:"out_trade_no,omitempty"`
	// 随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
	NonceStr string `xml:"nonce_str" url:"nonce_str,omitempty"`
	//时间戳	timestamp	String(10)	是	1412000000	时间戳，请见接口规则-参数规定
	Timestamp string `xml:"timestamp" url:"timestamp,omitempty"`
	// 签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	签名，详见签名生成算法
	Sign string `xml:"sign" url:"-"`
}

func (o *OrderQueryRequest) SetSign(info *AuthInfo) {
	o.Sign = info.MD5V(o)
}

type OrderQueryResponse struct {
	// 	返回状态码	return_code	是	String(16)	SUCCESS
	// SUCCESS/FAIL
	// 此字段是通信标识，非交易标识，交易是否成功需要查看trade_state来判断
	ReturnCode string `xml:"return_code" url:"return_code,omitempty"`
	// 返回信息	return_msg	否	String(128)	签名失败
	// 返回信息，如非空，为错误原因
	// 签名失败
	// 参数格式校验错误
	ReturnMsg string `xml:"return_msg" url:"return_msg,omitempty"`
	// 公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID
	AppID string `xml:"appid" url:"appid,omitempty"`
	// 商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
	MchID string `xml:"mch_id" url:"mch_id,omitempty"`
	// 随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
	NonceStr string `xml:"nonce_str" url:"nonce_str,omitempty"`
	// 设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号，
	DeviceInfo string `xml:"device_info" url:"device_info,omitempty"`
	// 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
	Sign string `xml:"sign" url:"-"`
	// 业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
	ResultCode string `xml:"result_code" url:"result_code,omitempty"`
	// 错误代码	err_code	否	String(32)	SYSTEMERROR	详细参见第6节错误列表
	ErrCode string `xml:"err_code" url:"err_code,omitempty"`
	// 错误代码描述	err_code_des	否	String(128)	系统错误	结果信息描述
	ErrCodeDes string `xml:"err_code_des" url:"err_code_des,omitempty"`
	// 用户标识	openid	是	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	用户在商户appid下的唯一标识
	Openid string `xml:"openid" url:"openid,omitempty"`
	// 是否关注公众账号	is_subscribe	否	String(1)	Y	用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	IsSubscribe string `xml:"is_subscribe" url:"is_subscribe,omitempty"`
	// 交易类型	trade_type	是	String(16)	JSAPI	调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	TradeType string `xml:"trade_type" url:"trade_type,omitempty"`
	// 交易状态	trade_state	是	String(32)	SUCCESS
	// SUCCESS—支付成功
	// REFUND—转入退款
	// NOTPAY—未支付
	// CLOSED—已关闭
	// REVOKED—已撤销（刷卡支付）
	// USERPAYING--用户支付中
	// PAYERROR--支付失败(其他原因，如银行返回失败)
	// 付款银行	bank_type	是	String(16)	CMC	银行类型，采用字符串类型的银行标识
	BankType string `xml:"bank_type" url:"bank_type,omitempty"`
	// 总金额	total_fee	是	Int	100	订单总金额，单位为分
	TotalFee string `xml:"total_fee" url:"total_fee,omitempty"`
	// 货币种类	fee_type	否	String(8)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	FeeType string `xml:"fee_type" url:"fee_type,omitempty"`
	// 现金支付金额	cash_fee	是	Int	100	现金支付金额订单现金支付金额，详见支付金额
	CashFee string `xml:"cash_fee" url:"cash_fee,omitempty"`
	// 现金支付货币类型	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType string `xml:"cash_fee_type" url:"cash_fee_type,omitempty"`
	// 代金券或立减优惠金额	coupon_fee	否	Int	100	“代金券或立减优惠”金额<=订单总金额，订单总金额-“代金券或立减优惠”金额=现金支付金额，详见支付金额
	CouponFee string `xml:"coupon_fee" url:"coupon_fee,omitempty"`
	// 代金券或立减优惠使用数量	coupon_count	否	Int	1	代金券或立减优惠使用数量
	CouponCount string `xml:"coupon_count" url:"coupon_count,omitempty"`
	// 代金券或立减优惠批次ID	coupon_batch_id_$n	否	String(20)	100	代金券或立减优惠批次ID ,$n为下标，从0开始编号
	CouponBatchId1 string `xml:"coupon_batch_id_1" url:"coupon_batch_id_1,omitempty"`
	// 代金券或立减优惠ID	coupon_id_$n	否	String(20)	10000 	代金券或立减优惠ID, $n为下标，从0开始编号
	CouponId1 string `xml:"coupon_id_" url:"coupon_id_,omitempty"`
	// 单个代金券或立减优惠支付金额	coupon_fee_$n	否	Int	100	单个代金券或立减优惠支付金额, $n为下标，从0开始编号
	CouponFee1 string `xml:"coupon_fee_1" url:"coupon_fee_1,omitempty"`
	// 微信支付订单号	transaction_id	是	String(32)	1009660380201506130728806387	微信支付订单号
	TransactionId string `xml:"transaction_id" url:"transaction_id,omitempty"`
	// 商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统的订单号，与请求一致。
	OutTradeNo string `xml:"out_trade_no" url:"out_trade_no,omitempty"`
	// 附加数据	attach	否	String(128)	深圳分店	附加数据，原样返回
	Attach string `xml:"attach" url:"attach,omitempty"`
	// 支付完成时间	time_end	是	String(14)	20141030133525	订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeEnd string `xml:"time_end" url:"time_end,omitempty"`
	// 交易状态描述	trade_state_desc	是	String(256)	支付失败，请重新下单支付	对当前查询订单状态的描述和下一步操作的指引
	TradeStateDesc string `xml:"trade_state_desc" url:"trade_state_desc,omitempty"`
}

func (o *OrderQueryResponse) VerifySign(info *AuthInfo, sign string) bool {
	if info.MD5V(o) == sign {
		return true
	}
	return false
}

type NotifyRequest struct {
	// 返回状态码	return_code	是	String(16)	SUCCESS
	// SUCCESS/FAIL
	// 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnCode string `xml:"return_code" url:"return_code,omitempty"`
	// 返回信息	return_msg	否	String(128)	签名失败
	// 返回信息，如非空，为错误原因
	// 签名失败
	// 参数格式校验错误
	ReturnMsg string `xml:"return_msg" url:"return_msg,omitempty"`
	// 公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
	AppID string `xml:"appid" url:"appid,omitempty"`
	// 商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
	MchID string `xml:"mch_id" url:"mch_id,omitempty"`
	// 设备号	device_info	否	String(32)	013467007045764	微信支付分配的终端设备号，
	DeviceInfo string `xml:"device_info" url:"device_info,omitempty"`
	// 随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位
	NonceStr string `xml:"nonce_str" url:"nonce_str,omitempty"`
	// 签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名算法
	Sign string `xml:"sign" url:"-"`
	// 业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
	ResultCode string `xml:"result_code" url:"result_code,omitempty"`
	// 错误代码	err_code	否	String(32)	SYSTEMERROR	错误返回的信息描述
	ErrCode string `xml:"err_code" url:"err_code,omitempty"`
	// 错误代码描述	err_code_des	否	String(128)	系统错误	错误返回的信息描述
	ErrCodeDes string `xml:"err_code_des" url:"err_code_des,omitempty"`
	// 用户标识	openid	是	String(128)	wxd930ea5d5a258f4f	用户在商户appid下的唯一标识
	Openid string `xml:"openid" url:"openid,omitempty"`
	// 是否关注公众账号	is_subscribe	否	String(1)	Y	用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	IsSubscribe string `xml:"is_subscribe" url:"is_subscribe,omitempty"`
	// 交易类型	trade_type	是	String(16)	JSAPI	JSAPI、NATIVE、APP
	TradeType string `xml:"trade_type" url:"trade_type,omitempty"`
	// 付款银行	bank_type	是	String(16)	CMC	银行类型，采用字符串类型的银行标识，银行类型见银行列表
	BankType string `xml:"bank_type" url:"bank_type,omitempty"`
	// 总金额	total_fee	是	Int	100	订单总金额，单位为分
	TotalFee string `xml:"total_fee" url:"total_fee,omitempty"`
	// 货币种类	fee_type	否	String(8)	CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	FeeType string `xml:"fee_type" url:"fee_type,omitempty"`
	// 现金支付金额	cash_fee	是	Int	100	现金支付金额订单现金支付金额，详见支付金额
	CashFee string `xml:"cash_fee" url:"cash_fee,omitempty"`
	// 现金支付货币类型	cash_fee_type	否	String(16)	CNY	货币类型，符合ISO4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType string `xml:"cash_fee_type" url:"cash_fee_type,omitempty"`
	// 代金券或立减优惠金额	coupon_fee	否	Int	10	代金券或立减优惠金额<=订单总金额，订单总金额-代金券或立减优惠金额=现金支付金额，详见支付金额
	CouponFee string `xml:"coupon_fee" url:"coupon_fee,omitempty"`
	// 代金券或立减优惠使用数量	coupon_count	否	Int	1	代金券或立减优惠使用数量
	CouponCount string `xml:"coupon_count" url:"coupon_count,omitempty"`
	// 代金券或立减优惠ID	coupon_id_$n	否	String(20)	10000	代金券或立减优惠ID,$n为下标，从0开始编号
	CouponId1 string `xml:"coupon_id_1" url:"coupon_id_1,omitempty"`
	// 单个代金券或立减优惠支付金额	coupon_fee_$n	否	Int	100	单个代金券或立减优惠支付金额,$n为下标，从0开始编号
	CouponFee1 string `xml:"coupon_fee_1" url:"coupon_fee_1,omitempty"`
	// 微信支付订单号	transaction_id	是	String(32)	1217752501201407033233368018	微信支付订单号
	TransactionId string `xml:"transaction_id" url:"transaction_id,omitempty"`
	// 商户订单号	out_trade_no	是	String(32)	1212321211201407033568112322	商户系统的订单号，与请求一致。
	OutTradeNo string `xml:"out_trade_no" url:"out_trade_no,omitempty"`
	// 商家数据包	attach	否	String(128)	123456	商家数据包，原样返回
	Attach string `xml:"attach" url:"attach,omitempty"`
	// 支付完成时间	time_end	是	String(14)	20141030133525	支付完成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeEnd string `xml:"time_end" url:"time_end,omitempty"`
}

func (o *NotifyRequest) VerifySign(info *AuthInfo, sign string) bool {
	if info.MD5V(o) == sign {
		return true
	}
	return false
}

type NotifyResponse struct {
	// 返回状态码	return_code	是	String(16)	SUCCESS
	// SUCCESS/FAIL
	// 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnCode string `xml:"return_code"`
	// 返回信息	return_msg	否	String(128)	签名失败
	// 返回信息，如非空，为错误原因
	// 签名失败
	// 参数格式校验错误
	ReturnMsg string `xml:"return_msg"`
}
