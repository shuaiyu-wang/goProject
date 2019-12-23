package test

import (
	"fmt"
)

func div(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到异常：%s\n", r)
		}
	}()

	if b < 0 {
		panic("除数需要大于0")
	}
	fmt.Println("余数为：", a/b)

}

func main() {
	// 捕捉内部的异常，程序并未终止，而是从异常中恢复，继续运行
	div(10, 0)
	// 捕捉主动的异常
	div(10, -1)
}
