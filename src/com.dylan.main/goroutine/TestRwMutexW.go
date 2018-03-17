package main

import (
	"fmt"
	"time"
	"sync"
)
var RWWmutex *sync.RWMutex;
var  RWWch chan int
// 读写模型2， 写的时候什么也做不了，
//也就是写保护
func main() {
	RWWmutex=new(sync.RWMutex)
	fmt.Println("start")
	RWWch=make(chan int )
	for i:=0; i<10; i++ {
		go TestRWWMutex(RWWch,i)
	}
	for i:=0; i<10; i++ {
		<-RWWch
	}
}

func TestRWWMutex(ch chan int,index int) {
	fmt.Println("进入写锁","index=",index)
	RWWmutex.Lock();

	//写东西..
	fmt.Println("正字写东西.....","index=",index)
	time.Sleep(1*time.Second)
	RWWmutex.Unlock()
	fmt.Println("离开写锁","index=",index)
	RWWch<-1
}
