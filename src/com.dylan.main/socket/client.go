package main

import (
	"net"
	"bytes"
	"io"
	"fmt"
)

func main() {

	add:=new(net.TCPAddr)
	add.IP=	net.ParseIP("123.207.215.205")
	add.Port=7777
	conn,err:=net.DialTCP("tcp",nil,add)
	if err!=nil{
		fmt.Println(err)
	}

	//conn.Write([]byte("hello"))
	conn.CloseWrite()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, conn)
	fmt.Println(string(buf.Bytes()))

}
