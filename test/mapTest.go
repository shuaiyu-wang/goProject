package test

import (
	"fmt"
	"strconv"
)

func main() {
	mymap := map[string]interface{}{
		"a": "a",
		"b": 1,
	}
	mymap2 := map[string]string{"a": "a", "b": ""}
	fmt.Println(mymap)
	fmt.Println(mymap2)

	//循环
	for key, value := range mymap {
		fmt.Println(key, value)
	}
	f := 0.01
	fmt.Println(strconv.FormatFloat(f*100, 'f', -1, 64))
}
