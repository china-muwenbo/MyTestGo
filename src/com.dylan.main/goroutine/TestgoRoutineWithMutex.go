package main

import (
	"sync"
	"fmt"
	"time"
)
//开启10个线程  同时去竞争一个互斥锁 谁有能力谁上

var mutex *sync.Mutex;
var  ch chan int


func main() {
	mutex=new(sync.Mutex)
	fmt.Println("start")
	ch=make(chan int )
	for i:=0; i<10; i++ {
		go TestMutex(ch,i)
	}
	for i:=0; i<10; i++ {
	<-ch
	}
}

func TestMutex(ch chan int,index int)  {
	fmt.Println("to enter mutex","index=",index)
	mutex.Lock();
	defer mutex.Unlock()
	defer fmt.Println("unLock","index=",index)
	fmt.Println("in mutex","index=",index)
	time.Sleep(2*time.Second)
	ch<-1
}

