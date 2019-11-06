package main

import (
	"fmt"
	)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	//ch <- "paul"
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}