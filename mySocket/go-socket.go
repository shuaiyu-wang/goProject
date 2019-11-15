package Mysocket

import (
	"encoding/json"
	"fmt"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

type user struct {
	RoomName string `json:"room_name"`
	Msg      string `json:"msg"`
}

func SocketMain() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		return
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		//s.SetContext("123")
		fmt.Println("connection:", s.ID())
		return nil
	})
	////create房间
	//server.onevent("/", "create", func(s socketio.conn, roomname string) {
	//	fmt.println("create room", s.id(), roomname)
	//	rooms:=s.rooms()
	//	rooms = append(rooms, roomname)
	//	fmt.println("创建"+roomname+"房间成功...", s.rooms())
	//})
	//加入房间
	server.OnEvent("/", "join", func(s socketio.Conn, roomName string) {
		fmt.Println("join room", s.ID(), roomName)
		s.Join(roomName)
		//给所用房间内的人发送消息，前端用hello事件接受
		server.BroadcastToRoom(roomName, "hello", "欢迎"+s.ID()+"来到"+roomName+"聊天室")
		//s.Emit("hello", "欢迎来到聊天室---"+s.ID())
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		//fmt.Println("notice:", msg, s.ID())
		////发送到聊天室
		//server.BroadcastToRoom("room1","reply", "["+s.ID()+"]------"+msg)
		//s.Emit("reply", "have"+msg)
	})
	server.OnEvent("/", "msg", func(s socketio.Conn, str string) {
		fmt.Println(str)
		paramMap := make(map[string]string)
		json.Unmarshal([]byte(str), &paramMap)
		fmt.Println(paramMap)
		//s.SetContext(msg)
		//s.Emit("hello","test"+msg)
		server.BroadcastToRoom(paramMap["room_name"], "hello", "["+s.ID()+"]------"+paramMap["msg"])
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})
	//server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
	//	fmt.Println("closed", reason)
	//})
	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Print("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
