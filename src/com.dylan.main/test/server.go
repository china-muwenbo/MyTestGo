package main

import (
	"net"
	"fmt"
	"time"
	"bytes"
	"io"
)

func main() {
	add:=new(net.TCPAddr)
	add.IP= net.ParseIP("127.0.0.1")
	add.Port=7777
	tcpl,err:=net.ListenTCP("tcp", add)
	if err != nil {
		fmt.Println(err)
	}
	for{
		conn,err:=tcpl.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//handleClient(conn)
		go Read1(conn)
	}
}

func Read1(conn net.Conn){
	buf := make([]byte, 128)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取服务器数据异常:", err.Error())
	}
	fmt.Println("read",string(buf))

}



func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)
	if err != nil {
		fmt.Println("读取错误")
	}
	fmt.Println(string(buf.Bytes()))
	data:=buf.Bytes()
	fmt.Println(len(data))
	if len(data)==0 {
		daytime := time.Now().String()
		//不需要关心返回值
		conn.Write([]byte(daytime))
	}else{
		//不需要关心返回值
		conn.Write([]byte(string(buf.Bytes())))
	}
}