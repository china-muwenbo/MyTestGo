package main

import (
	"fmt"
	"time"
	"sync"
)

//读写锁模型一
// 10 协程一起读


var RWmutex *sync.RWMutex;
var  RWch chan int

func main() {
	RWmutex=new(sync.RWMutex)
	fmt.Println("start")
	RWch=make(chan int )
	for i:=0; i<10; i++ {
		go TestRWMutex(i)
	}
	for i:=0; i<10; i++ {
		<-RWch
	}



}

func TestRWMutex(index int)  {
	fmt.Println("进入读写锁，准备读点东西","index=",index)
	RWmutex.RLock();
	//读取数据
	fmt.Println("读点东西..","index=",index)
	time.Sleep(4*time.Second)
	RWmutex.RUnlock()
	fmt.Println("离开读写锁..","index=",index)
	RWch<-1
}

