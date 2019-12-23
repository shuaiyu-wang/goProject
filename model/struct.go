package model

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Id   int
	Name string
	Sex  bool
}

func main() {
	stu := student{
		Id:   1,
		Name: "wsy",
		Sex:  false,
	}
	json1, _ := json.Marshal(stu)
	fmt.Println(string(json1), 1)
	var stu2 student
	//反序列化
	json.Unmarshal([]byte(json1), &stu2)
	fmt.Println(stu2, 2)

	//定义map
	maps := make(map[string]interface{})
	//定义数组
	arr := make([]map[string]interface{}, 0)
	maps["id"] = 2
	maps["name"] = "wsy2"
	maps["sex"] = true
	maps["id"] = 3333
	arr = append(arr, maps)
	maps = make(map[string]interface{})

	maps["id"] = 3
	maps["name"] = "wsy3"
	maps["sex"] = true
	maps["id"] = 3334
	arr = append(arr, maps)
	json2, err := json.Marshal(arr)
	fmt.Println(string(json2), err)
}
