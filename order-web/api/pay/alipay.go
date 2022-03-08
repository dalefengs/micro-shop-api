package pay

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"go.uber.org/zap"
	"micro-shop-api/order-web/global"
	"micro-shop-api/order-web/global/response"
	status2 "micro-shop-api/order-web/global/status"
	"micro-shop-api/order-web/proto"
	"net/http"
)

var (
	TradeStatusWaitBuyerPay string = "WAIT_BUYER_PAY" // 交易创建等待买家支付
	TradeStatusClose               = "TRADE_CLOSED"   // 交易关闭
	TradeStatusSuccess             = "TRADE_SUCCESS"  // 支付成功
	TradeStatusFinished            = "TRADE_FINISHED" // 交易完结(不可退款）
)

// Notify 支付宝异步通知
func Notify(ctx *gin.Context) {

	// 解析异步请求的参数
	notifyReq, err := alipay.ParseNotifyToBodyMap(ctx.Request)
	if err != nil {
		zap.S().Errorf("解析参数错误:%s", err.Error())
		response.Fail(ctx, status2.InvalidParameter)
		return
	}
	// 验证消息是否是支付宝发送的 保证安全性
	// 支付宝步通知验签（公钥证书模式）
	ok, err := alipay.VerifySignWithCert(global.Config.Alipay.AlipayPublicContentRSA2, notifyReq)
	if !ok || err != nil {
		zap.S().Errorf("支付宝异步通知验签失败:%s", err.Error())
		response.Fail(ctx, status2.InnerError)
		return
	}
	// 以下是业务逻辑
	// 更新订单状态
	orderSn := notifyReq["out_trade_no"].(string)
	orderStatus := notifyReq["trade_status"].(string)
	//payTime := notifyReq["gmt_payment"].(string)
	_, err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{OrderSn: orderSn, Status: orderStatus})
	if err != nil {
		zap.S().Errorf("支付成功但更新状态失败,支付状态:%s, 订单号:%s, err:%s", orderStatus, orderSn, err.Error())
		response.Fail(ctx, status2.InnerError)
		return
	}

	// ====异步通知，返回支付宝平台的信息====
	//    文档：https://opendocs.alipay.com/open/203/105286
	//    程序执行完后必须打印输出“success”（不包含引号）。如果商户反馈给支付宝的字符不是success这7个字符，支付宝服务器会不断重发通知，直到超过24小时22分钟。一般情况下，25小时以内完成8次通知（通知的间隔频率一般是：4m,10m,10m,1h,2h,6h,15h）
	// 	  也就是最大努力通知
	ctx.String(http.StatusOK, "success")
}

type Alipay struct {
	client alipay.Client
}

func NewAlipay() (*Alipay, error) {
	cert := global.Config.Alipay
	// 初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(cert.Appid, cert.AppPrivateKey, false)
	if err != nil {
		return nil, err
	}
	// 打开Debug开关，输出日志，默认关闭
	//client.DebugSwitch = gopay.DebugOn

	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).      // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).     // 设置签名类型，不设置默认 RSA2
							SetReturnUrl(cert.ReturnUrl). // 设置返回URL，付款结束后跳转的url
							SetNotifyUrl(cert.NotifyUrl)  // 设置异步通知URL

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign(cert.AlipayPublicContentRSA2)

	// 证书内容
	err = client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)

	if err != nil {
		return nil, err
	}
	return &Alipay{*client}, nil
}

// GetAlipayUrl 获取支付 url
// subject 标题
// 订单号，支付成功后会返回
// 订单金额
func (c *Alipay) GetAlipayUrl(subject, outTradeNo, amount string) (string, error) {

	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", subject). // 标题
					Set("out_trade_no", outTradeNo).                             // 订单号，支付成功后会返回
					Set("total_amount", amount).                                 // 订单金额
					Set("timeout_express", global.Config.Alipay.TimeoutExpress). // 支付超时时间
					Set("product_code", "FAST_INSTANT_TRADE_PAY")                // 必填 具体参考文档

	payUrl, err := c.client.TradePagePay(context.Background(), bm)
	if err != nil {
		return "", err
	}
	return payUrl, nil
}
