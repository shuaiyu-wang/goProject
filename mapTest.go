package main

import "fmt"

func main() {
	mymap := map[string]interface{}{
		"a":"a",
		"b":1,
	}
	mymap2 := map[string]string{"a":"a","b":""}
	fmt.Println(mymap)
	fmt.Println(mymap2)
}