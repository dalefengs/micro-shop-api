package goods

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"micro-shop-api/goods-web/global"
	"micro-shop-api/goods-web/proto"
	"net/http"
	"strconv"

	"micro-shop-api/goods-web/global/response"
	status2 "micro-shop-api/goods-web/global/status"
)

// HandlerGrpcErrorToHttp grpc异常转换为http异常
func HandlerGrpcErrorToHttp(err error, c *gin.Context) {
	if err == nil {
		return
	}
	s, ok := status.FromError(err)
	if !ok {
		return
	}
	var (
		msg      string
		httpCode int
		code     = status2.ServerError
	)
	switch s.Code() {
	case codes.NotFound:
		httpCode = http.StatusNotFound
		msg = s.Message()
	case codes.Internal:
		httpCode = http.StatusInternalServerError
		msg = "内部错误"
	case codes.InvalidArgument:
		httpCode = http.StatusBadRequest
		code = status2.InvalidParameter
	case codes.Unavailable:
		httpCode = http.StatusInternalServerError
		code = status2.UnavailableServer
	default:
		httpCode = http.StatusInternalServerError
	}
	response.FailResponse(c, httpCode, code, msg, "")
}

func List(c *gin.Context) {
	request := proto.GoodsFilterRequest{}

	priceMin, _ := strconv.Atoi(c.DefaultQuery("pmin", "0"))
	request.PriceMin = int32(priceMin)

	priceMax, _ := strconv.Atoi(c.DefaultQuery("pmax", "0"))
	if priceMax > 0 {
		request.PriceMax = int32(priceMax)
	}

	// 是否热门

	if isHot := c.DefaultQuery("ih", "0"); isHot == "1" {
		request.IsHot = true
	}

	// 是否最新
	if isNew := c.DefaultQuery("in", "0"); isNew == "1" {
		request.IsNew = true
	}

	// 是否是Tab
	if isTab := c.DefaultQuery("it", "0"); isTab == "1" {
		request.IsTab = true
	}

	if category_id, _ := strconv.Atoi(c.DefaultQuery("c", "0")); category_id != 0 {
		request.TopCategory = int32(category_id)
	}
	pn, _ := strconv.Atoi(c.DefaultQuery("pn", "1"))
	pnum, _ := strconv.Atoi(c.DefaultQuery("pnum", "10"))
	keyword := c.DefaultQuery("q", "")
	if brand_id, _ := strconv.Atoi(c.DefaultQuery("c", "0")); brand_id != 0 {
		request.Brand = int32(brand_id)
	}

	request.Pages = int32(pn)
	request.PagePerNums = int32(pnum)
	request.KeyWords = keyword

	r, err := global.GoodsSrvClient.GoodsList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 [用户列表] 失败",
			"msg", err.Error())
		HandlerGrpcErrorToHttp(err, c)
		return
	}
	reMap := map[string]interface{}{
		"total": r.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, value := range r.Data {
		goodsList = append(goodsList, map[string]interface{}{
			"id":          value.Id,
			"name":        value.Name,
			"goods_brief": value.GoodsBrief,
			"desc":        value.GoodsDesc,
			"ship_free":   value.ShipFree,
			"images":      value.Images,
			"desc_images": value.DescImages,
			"front_image": value.GoodsFrontImage,
			"shop_price":  value.ShopPrice,
			"ctegory": map[string]interface{}{
				"id":   value.Category.Id,
				"name": value.Category.Name,
			},
			"brand": map[string]interface{}{
				"id":   value.Brand.Id,
				"name": value.Brand.Name,
				"logo": value.Brand.Logo,
			},
			"is_hot":  value.IsHot,
			"is_new":  value.IsNew,
			"on_sale": value.OnSale,
		})
	}
	reMap["data"] = goodsList

	fmt.Println(reMap)
	response.SuccessData(c, reMap)

}
