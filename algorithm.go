package main

import "fmt"

func main() {
	arr := []int{1,89,64,9,101,77,23}
	//1,64,9,89,77,23,101
	//1,9,64,77,23,89,101
	//1,9,64,23,77,89,101
	//1,9,23,64,77...
	for i:=0;i<len(arr);i++{
		for j:= 0;j<len(arr)-i-1;j++{
			if(arr[j]>arr[j+1]){
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
		fmt.Println(arr,i)
	}
	fmt.Println(arr)
}