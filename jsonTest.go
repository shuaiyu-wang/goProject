package main

import (
	"encoding/json"
	"fmt"
	"ytherp/common/utils"
)

type Test struct {
	name string `json:"keyword" valid:"MaxSize(100)" vdesc:"模糊查询关键词应在100位以内" description:"模糊匹配关键字"`
}

func main() {
	//var test Test
	test01 := Test{
		"wsy",
	}
	fmt.Println(test01)

	var b []byte
	json.Unmarshal(b, &test01)
	mymap := utils.StructToMap(&test01)
	fmt.Println(mymap)

	str := ",..."
	str += "xxx"
	fmt.Println(str)

	//utils.GetTimeFormDate()

	var flo float64
	flo = 1.41
	//flo = fmt.Sprintf("%.0f", flo)
	fmt.Printf("%.0f", flo)
}