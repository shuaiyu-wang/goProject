package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//time to string
	t1 := time.Now().Format("2006-01-02")
	t1 = strings.Replace(time.Now().Format("20060102150405.000"), ".", "", 1)
	ran := fmt.Sprintf("%05v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000))
	//layout := strings.Replace("yyyy/MM/dd HH:mm:ss.SSS", "SSS", "000",-1)
	fmt.Println("ti", t1, ran)

	//string to time
	time2, _ := time.Parse("2006-01-02", t1)
	time3, _ := time.ParseInLocation("2006-01-02", t1, time.Local)
	fmt.Println(time2)
	fmt.Println(time3)
	t2 := time2.Format("2006-01-02")
	t3 := time3.Format("2006-01-02")
	fmt.Println(t2)
	fmt.Println(t3)

	date, _ := time.Parse("2006-01-02 15:04:05", "2019-07-05 10:09:05")
	fmt.Println(date)

	fmt.Println("........", time.Now().Year())

	a := []int{0, 1, 2, 3, 4}
	//删除第i个元素
	i := 2
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)
	fmt.Println("随机数", CreateNonceStr())
}

func CreateNonceStr() string {
	str := ""
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 32; i++ {
		str += SubString(chars, rand.Intn(len(chars)), 1)
	}
	return str
}

func SubString(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
