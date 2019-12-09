package hprose

import (
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
	"strconv"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

func hello(name string) string {
	return "Hello " + name + "!"
}

func add(user *User) string {
	return "user.id " + strconv.Itoa(user.Id) + "!"
}

func ServerMain() {
	service := rpc.NewHTTPService()
	service.AddFunction("add", add, rpc.Options{})
	http.ListenAndServe(":2000", service)
}
