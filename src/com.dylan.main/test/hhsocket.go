package main

import (
	"fmt"
	"net"

)

func main() {
	service := "127.0.0.1:7777"
	//service := "192.168.1.13:7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	conn,err:= net.DialTCP("tcp", nil, tcpAddr)
	if err 	!= nil{
		fmt.Println(err)
	}
	//conn.Write([]byte("hello"))
	//conn.CloseWrite();

	go read(conn)

	sendMsg := ""
	for{
		_, err := fmt.Scanln(&sendMsg)
		fmt.Println(sendMsg)
		if err != nil {
			fmt.Println(err)
		}
		_, err = conn.Write([]byte(sendMsg))

		if err !=nil{
			fmt.Println(err)
		}
	}
}

func read(conn net.Conn){
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取服务器数据异常:", err.Error())
		}
		fmt.Println("read",string(buf))
}

