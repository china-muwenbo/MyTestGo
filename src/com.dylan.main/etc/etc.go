package main

import (
	"fmt"
	"time"
)
var ch chan  int
func main() {

	ch =make(chan  int)
	for i:=0;i<10 ;i++  {
		go test(i)
	}

	for i:=0;i<10 ;i++  {
		<-ch
	}

}
func test(index int ){
	time.Sleep(time.Second*1)
	defer fmt.Println("退出test ",index)
	ch<-index
}

