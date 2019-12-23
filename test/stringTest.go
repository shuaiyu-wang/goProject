package test

import (
	"fmt"
	"strconv"
)

type myString string
type myInt int

func main() {
	str := "hello world"
	i := 1
	fmt.Println(str + strconv.Itoa(i+1) + "12")
	var char byte = 'a'
	for index, s := range str {

		fmt.Println(index, string(s), len(str), char)
	}

	fmt.Println(myString("byte"), myInt(23))
	fmt.Println(myString("x"))

	s := []string{"x1", "x2", "x3"}
	for x, v := range s {
		//fmt.Print(strconv.Itoa(x)+",")
		if x == 1 {
			v = "v2"
		} else {
			v = v
		}
		fmt.Print(v + ",")
	}
}
