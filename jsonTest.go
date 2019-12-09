package main

import (
	"encoding/json"
	"fmt"
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

	//var b []byte
	//json.Unmarshal(b, &test01)
	//mymap := utils.StructToMap(&test01)
	//fmt.Println(mymap)

	str := ",..."
	str += "xxx"
	fmt.Println(str)

	//utils.GetTimeFormDate()

	var flo float64
	flo = 1.41
	//flo = fmt.Sprintf("%.0f", flo)
	fmt.Printf("%.0f", flo)

	str1 := "[{\"id\":\"1752\",\"name\":\"采购分仓\"},{\"id\":\"1753\",\"name\": \"采购分仓\"}]"

	paramMap := make([]map[string]string, 0)
	json.Unmarshal([]byte(str1), &paramMap)
	fmt.Println(paramMap)
}
