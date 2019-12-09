package hprose

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
)

type Stub struct {
	Hello      func(string) (string, error)
	AsyncHello func(func(string, error), string) `name:"hello"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

type UserService struct {
	Add func(user *User) (id string, err error)
}

func ClientMain() {
	client := rpc.NewClient("http://127.0.0.1:2000/")
	var stub *Stub
	client.UseService(&stub)
	stub.AsyncHello(func(result string, err error) {
		fmt.Println(result, err)
	}, "async world")
	//调用远程的hello
	fmt.Println(stub.Hello("world"))
}

func MyNewService() {
	client := rpc.NewClient("http://127.0.0.1:2000/")
	var userService *UserService
	client.UseService(&userService)
	user := new(User)
	user.Id = 1
	fmt.Println(userService.Add(user))
}
