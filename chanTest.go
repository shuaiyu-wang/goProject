package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	done2 := make(chan bool)
	go helloChan(done)
	go helloChan2(done2)
	<- done
	<- done2
	fmt.Println("main function")

}
func helloChan(done chan bool){
	for i:=0;i<=5;i++{
		time.Sleep(250*time.Millisecond)
		fmt.Println("hello chan",i)
	}
	done<-true
}
func helloChan2(done2 chan bool){
	for i:='a';i<='e';i++{
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("hello chan2 %c\n",i)
	}
	done2<-true
}