package main

import (
		"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//time to string
	t1 := time.Now().Format("2006-01-02")
	fmt.Println(t1)

	//string to time
	time2,_ := time.Parse("2006-01-02", t1)
	time3, _ := time.ParseInLocation("2006-01-02", t1, time.Local)
	fmt.Println(time2)
	fmt.Println(time3)
	t2 := time2.Format("2006-01-02")
	t3 := time3.Format("2006-01-02")
	fmt.Println(t2)
	fmt.Println(t3)

	date,_ := time.Parse("2006-01-02 15:04:05","2019-07-05 10:09:05")
	fmt.Println(date)

	fmt.Println("........",time.Now().Year())

	a := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
}
