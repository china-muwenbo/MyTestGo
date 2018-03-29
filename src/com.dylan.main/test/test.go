package main

import (
	"bytes"
	"fmt"
	"time"
	"net"
	"io"
)
type HmConn struct {
	conns []net.Conn
}

func main() {

	hmConn:=new(HmConn)
	hmConn.conns= make([]net.Conn,100)

	add:=new(net.TCPAddr)

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
		go handleClient(conn)
	}
}

func(  hm *HmConn) handleClient(conn net.Conn) {
	//defer conn.Close()
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

func( hm * HmConn) send(msg []byte){

	for _,con := range hm.conns{
		con.Write(msg)
	}
}
