package test

import "fmt"

func main() {
	arrs := make([]AllGoodsVo2, 0)
	var arr AllGoodsVo2
	arr.GoodsCode = "0"
	arrs = append(arrs, arr)
	arr.GoodsCode = "1"
	arrs = append(arrs, arr)
	arr.GoodsCode = "2"
	arrs = append(arrs, arr)
	for i, _ := range arrs {
		if arrs[i].GoodsCode == "0" {
			arrs[i].GoodsCode = "x"
		}
	}

	fmt.Println(arrs[0].GoodsCode)
}

type AllGoodsVo2 struct {
	RowNum             int    `json:"row_num" description:"序号"`
	Id                 string `json:"id" description:"商品id"`
	Name               string `json:"name" description:"商品名称"`
	CateLv1Id          string `json:"cate_lv1_id" description:"一级类目ID"`
	CateLv1Name        string `json:"cate_lv1_name" description:"一级类目名称"`
	CateLv2Id          string `json:"cate_lv2_id" description:"二级类目ID"`
	CateLv2Name        string `json:"cate_lv2_name" description:"二级类目名称"`
	PrdId              string `json:"prd_id" description:"产品ID"`
	PrdName            string `json:"prd_name" description:"产品名称"`
	Unit               string `json:"unit" description:"单位（编辑时使用）"`
	UnitId             string `json:"unit_id" description:"基本单位Id"`
	UnitName           string `json:"unit_name" description:"基本单位名称"`
	RatUnitId          string `json:"rat_unit_id" description:"转化单位Id"`
	RatUnitName        string `json:"rat_unit_name" description:"转化单位名称"`
	RequireUnitId      string `json:"require_unit_id" description:"转化单位Id"`
	RequireUnitName    string `json:"require_unit_name" description:"商品基本单位名称"`
	Memo               string `json:"memo" description:"备注"`
	GoodsCode          string `json:"goods_code" description:"商品编号"`
	MainPicUrl         string `json:"main_pic_url" valid:"Required" vdesc:"主图片地址不能为空" required:"true" description:"主图片"`
	CategoryId         string `json:"category_id" description:"商品属性Id"`
	CategoryName       string `json:"category_name" description:"商品属性名称"`
	SpecificationId    string `json:"specification_id" description:"商品规格Id"`
	SpecificationName  string `json:"specification_name" description:"商品规格名称"`
	GoodsId            string `json:"goods_id" description:"商品Id"`
	Rate               int    `json:"rate" description:"倍数"`
	IsRatingDesc       string `json:"is_rating_desc" description:"是否为评价表商品"`
	RatingFormDetailId string `json:"rating_form_detail_id" description:"评价表详情ID"`
	GoodsName          string `json:"goods_name" description:"商品名称"`
	Recipient          string `json:"recipient" description:"领料人"`
	AgentId            string `json:"agent_id"    description:"代理商Id"`
	AgentCustPriceId   string `json:"agent_cust_price_id" description:"外销价id"`
}
