package test

import "fmt"

//func main() {
//	arr := make([]int,0)
//	arr = append(arr, 4)
//	//var err error
//	//fmt.Println(1,arr)
//	//change(arr)
//	//fmt.Println(3,arr)
//
//
//	//slic := arr[:5]
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//	arr = append(arr, 4)
//	fmt.Println(cap(arr))
//
//}
//func change(arr [3]int){
//	arr[0] = 2
//	fmt.Println(2,arr)
//}

//func changeLocal(num [5]int) {
//	num[0] = 55
//	fmt.Println("inside function ", num)
//}
//func main() {
//	num := [...]int{5, 6, 7, 8, 8}
//	fmt.Println("before passing to function ", num)
//	changeLocal(num) //num is passed by value
//	fmt.Println("after passing to function ", num)
//}

func change(s ...string) {
	//s[0] = "Go"
	fmt.Println(s)
	s = append(s, "playground")
	fmt.Println(s)
}

func main() {
	//welcome := []string{"hello", "world"}
	//change(welcome...)
	//fmt.Println(welcome)
	arr := make([]string, 4, 4)
	change(arr...)
	fmt.Println(arr)
}
