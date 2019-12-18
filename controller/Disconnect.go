package controller

import (
	"HelloWorld/io/network/route"
	"HelloWorldServer/model/Login"
	"HelloWorldServer/packet"
	"fmt"
)

func init() {
	route.Register(packet.Disconnect{}, Disconnect)
}

func Disconnect(ID uint64) {
	fmt.Println("一个连接断开:", ID)
	Login.SignOut(ID)
}
