package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
)

type students struct {
	Name string ``
}

func main() {
	//data,err := ioutil.ReadFile("D:\\ioTest.txt")
	//fmt.Println(string(data),err)

	//file,_ := os.Open("D:\\ioTest.txt")
	//d:=bufio.NewScanner(file)
	//for d.Scan(){
	//	fmt.Println(d.Text())
	//	break
	//}
	//file.Close()

	//f, err := os.Create("test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	_, err := os.OpenFile("/tmp/test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	//newLine := "File handling is easy."
	////_, err = fmt.Fprintln(f, newLine)
	//_, err = fmt.Fprintln(f, newLine)
	//if err != nil {
	//	fmt.Println(err)
	//	f.Close()
	//	return
	//}
	//err = f.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("file appended successfully")
	//fmt.Println("xxx",new(students))
	//fmt.Println("xxx",students{Name:"wsy"}.Name)
	//
	//fmt.Println("xxx", New("wsy2"))
	//
	//i := new(int)
	//fmt.Println(*i)
	//b := new(bool)
	//fmt.Println(*b)
	//
	//var a interface{}
	//reflect.ValueOf(a)

}

func New(name string) students {
	return students{Name: name}
}
