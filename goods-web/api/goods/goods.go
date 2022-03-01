package goods

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"micro-shop-api/goods-web/api"
	"micro-shop-api/goods-web/global"
	"micro-shop-api/goods-web/proto"
	"micro-shop-api/goods-web/validator/forms"
	"strconv"

	"micro-shop-api/goods-web/global/response"
	status2 "micro-shop-api/goods-web/global/status"
)

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
		api.HandlerGrpcErrorToHttp(err, c)
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

func New(c *gin.Context) {
	goodsForm := forms.GoodsForm{}
	err := c.ShouldBindJSON(&goodsForm)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	goodsClient := global.GoodsSrvClient

	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	response.SuccessMsgData(c, "新建商品成功", rsp)
}

func Detail(c *gin.Context) {
	id := c.Param("id")
	goodsId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		response.Fail(c, status2.GoodsNotFind)
		return
	}
	r, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{Id: int32(goodsId)})
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	rsp := map[string]interface{}{
		"id":          r.Id,
		"name":        r.Name,
		"goods_brief": r.GoodsBrief,
		"desc":        r.GoodsDesc,
		"ship_free":   r.ShipFree,
		"images":      r.Images,
		"desc_images": r.DescImages,
		"front_image": r.GoodsFrontImage,
		"shop_price":  r.ShopPrice,
		"ctegory": map[string]interface{}{
			"id":   r.Category.Id,
			"name": r.Category.Name,
		},
		"brand": map[string]interface{}{
			"id":   r.Brand.Id,
			"name": r.Brand.Name,
			"logo": r.Brand.Logo,
		},
		"is_hot":  r.IsHot,
		"is_new":  r.IsNew,
		"on_sale": r.OnSale,
	}
	response.SuccessData(c, rsp)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	goodsId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		response.Fail(c, status2.GoodsNotFind)
		return
	}
	_, err = global.GoodsSrvClient.DeleteGoods(c, &proto.DeleteGoodsInfo{
		Id: int32(goodsId),
	})
	if err != nil {
		response.FailCodeMsg(c, status2.InvalidOperation, "删除失败")
		return
	}
	response.SuccessMsg(c, "删除成功")

}

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	goodsId, err := strconv.ParseInt(id, 10, 32)
	statusForm := forms.GoodsStatusForm{}
	err = c.ShouldBindJSON(&statusForm)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	_, err = global.GoodsSrvClient.UpdateGoods(c, &proto.CreateGoodsInfo{
		Id:     int32(goodsId),
		IsNew:  *statusForm.IsNew,
		IsHot:  *statusForm.IsHot,
		OnSale: *statusForm.OnSale,
	})

	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	response.SuccessMsg(c, "更新成功")
}

func Update(c *gin.Context) {
	id := c.Param("id")
	goodsId, err := strconv.ParseInt(id, 10, 32)
	goodsForm := forms.GoodsForm{}
	err = c.ShouldBindJSON(&goodsForm)
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	goodsClient := global.GoodsSrvClient

	rsp, err := goodsClient.CreateGoods(context.Background(), &proto.CreateGoodsInfo{
		Id:              int32(goodsId),
		Name:            goodsForm.Name,
		GoodsSn:         goodsForm.GoodsSn,
		Stocks:          goodsForm.Stocks,
		MarketPrice:     goodsForm.MarketPrice,
		ShopPrice:       goodsForm.ShopPrice,
		GoodsBrief:      goodsForm.GoodsBrief,
		ShipFree:        *goodsForm.ShipFree,
		Images:          goodsForm.Images,
		DescImages:      goodsForm.DescImages,
		GoodsFrontImage: goodsForm.FrontImage,
		CategoryId:      goodsForm.CategoryId,
		BrandId:         goodsForm.Brand,
	})
	if err != nil {
		api.HandlerGrpcErrorToHttp(err, c)
		return
	}
	response.SuccessMsgData(c, "更新成功", rsp)
}
