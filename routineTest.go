package main

import (
	"fmt"
	"time"
)

func main() {

	go helloRoutine()
	go helloRoutine2()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

func helloRoutine(){
	for i:=0;i<=5;i++{
		time.Sleep(250*time.Millisecond)
		fmt.Println("hello routine",i)
	}
}

func helloRoutine2(){
	for i:='a';i<='e';i++{
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("hello routine2 %c\n",i)
	}
}