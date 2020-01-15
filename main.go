package main

import "fmt"

func main() {
	//var i *int
	//a := 3
	////var name string = "wsy"
	//i = &a
	//a = *i
	//const CON = "con"
	////fmt.Println(i)
	//fmt.Printf("hello go %v , %v\n", i, a)
	//
	//var j = 1
	//var m int
	//m = 2
	//var n int = 3
	//fmt.Printf("%v,%v,%v\n", j, m, n)
	//
	//fmt.Println(Books{title: "aaa", author: "bbb", subject: "ccc", book_id: 1})
	//
	//mymap := make(map[string][]string)
	//myarr := []string{"1", "2", "3"}
	//mymap["a"] = myarr
	//fmt.Printf("%v,%v\n", mymap, myarr)
	//zap.Init()
	//zap.Logger.Info("zap------------test")
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("..................", info)
		}
	}()
	test(0)
}

func test(num int) {
	fmt.Println(1 / num)
}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}
