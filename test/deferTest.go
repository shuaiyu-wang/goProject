package test

import (
	"errors"
	"fmt"
)

func main() {
	//i := 0
	//defer fmt.Println(i)
	//i ++
	//defer fmt.Println(i)

	//foo()
	//foo()
	//var i int
	//var err error
	//i,err= c()
	//fmt.Println(i,err)
	a()
}

func c() (i int, err error) {
	defer func() {
		//start := time.Now()
		//fmt.Println(i);
		i++
		fmt.Println(i)
	}()
	//time.Sleep(5*time.Second)
	return 1, errors.New("xxx")
}

//func trace(funcName string) func(){
//	start := time.Now()
//	fmt.Printf("function %s enter\n",funcName)
//	return func(){
//		fmt.Printf("function %s exit (elapsed %s)",funcName,time.Since(start))
//	}
//}
//
//func foo(){
//	defer trace("foo()")()
//	time.Sleep(5*time.Second)
//}

func a() {
	i := 0
	i++
	//defer 执行阶段处于 return 之后，函数返回之前
	defer fmt.Println("defer", i) //那这个为什么输出为 1 ？不应该是 2 吗
	i++
	fmt.Println("i=", i)
	return
}
