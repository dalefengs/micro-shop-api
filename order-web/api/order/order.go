package order

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"micro-shop-api/order-web/api"
	"micro-shop-api/order-web/api/pay"
	"micro-shop-api/order-web/global"
	"micro-shop-api/order-web/global/response"
	status2 "micro-shop-api/order-web/global/status"
	"micro-shop-api/order-web/proto"
	"micro-shop-api/order-web/validator/forms"
	"strconv"
)

func List(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	pageStr := ctx.DefaultQuery("page", "1")
	limitSrt := ctx.DefaultQuery("limit", "10")
	// 因为数字有默认的值 所以不需要处理

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitSrt)
	resuest := proto.OrderFilterRequest{
		UserId:      int32(userId.(uint)),
		Pages:       int32(page),
		PagePerNums: int32(limit),
	}

	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &resuest)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	result := map[string]interface{}{
		"total": rsp.Total,
	}

	orderList := make([]interface{}, rsp.Total)
	for _, item := range rsp.Data {
		data := map[string]interface{}{
			"id":       item.Id,
			"user_id":  item.UserId,
			"order_sn": item.OrderSn,
			"pay_type": item.PayType,
			"status":   item.Status,
			"post":     item.Post,
			"address":  item.Address,
			"name":     item.Name,
			"mobile":   item.Mobile,
		}
		orderList = append(orderList, data)
	}

	result["data"] = orderList
	response.SuccessData(ctx, result)

}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	orderId, _ := strconv.Atoi(id)
	request := proto.OrderRequest{
		Id: int32(orderId),
	}
	rsp, err := global.OrderSrvClient.OrderDetail(context.Background(), &request)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, ctx)
		return
	}
	result := gin.H{}
	result["id"] = rsp.OrderInfo.Id
	result["status"] = rsp.OrderInfo.Status
	result["user"] = rsp.OrderInfo.UserId
	result["post"] = rsp.OrderInfo.Post
	result["total"] = rsp.OrderInfo.Total
	result["address"] = rsp.OrderInfo.Address
	result["name"] = rsp.OrderInfo.Name
	result["mobile"] = rsp.OrderInfo.Mobile
	result["pay_type"] = rsp.OrderInfo.PayType
	result["order_sn"] = rsp.OrderInfo.OrderSn

	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		tmpMap := gin.H{
			"id":    item.GoodsId,
			"name":  item.GoodsName,
			"image": item.GoodsImage,
			"price": item.GoodsPrice,
			"nums":  item.Nums,
		}

		goodsList = append(goodsList, tmpMap)
	}
	result["goods"] = goodsList
	client, err := pay.NewAlipay()
	if err != nil {
		zap.S().Infof("调用支付宝支付失败:%s", err.Error())
		response.FailCodeMsg(ctx, status2.InnerError, "调用支付宝支付失败")
		return
	}
	payUrl, err := client.GetAlipayUrl("订单支付", rsp.OrderInfo.OrderSn, strconv.FormatFloat(float64(rsp.OrderInfo.Total), 'f', 2, 32))
	if err != nil {
		zap.S().Infof("生成支付宝url:%s", err.Error())
		response.FailCodeMsg(ctx, status2.InnerError, "支付失败")
		return
	}
	result["alipayUrl"] = payUrl
	response.SuccessData(ctx, result)
}

func New(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	orderForm := forms.OrderForm{}
	err := ctx.ShouldBindJSON(&orderForm)
	if err != nil {
		api.HandleValidatorError(err, ctx)
		return
	}

	request := proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Address: orderForm.Address,
		Mobile:  orderForm.Mobile,
		Name:    orderForm.Name,
		Post:    orderForm.Post,
	}
	rsp, err := global.OrderSrvClient.CreateOrder(context.Background(), &request)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, ctx)
		return
	}

	client, err := pay.NewAlipay()
	if err != nil {
		response.FailCodeMsg(ctx, status2.InnerError, "创建支付失败")
		return
	}
	url, err := client.GetAlipayUrl("订单支付", rsp.OrderSn, strconv.FormatFloat(float64(rsp.Total), 'f', 5, 32))
	if err != nil {
		response.FailCodeMsg(ctx, status2.InnerError, "支付调用失败")
		return
	}

	result := gin.H{
		"id":         rsp.Id,
		"order_sn":   rsp.OrderSn,
		"alipay_url": url,
	}

	response.SuccessData(ctx, result)
}
