package main

import (
	"sync"
	"fmt"
	"time"
)

//读写同时发生的模型
//读的时候的写 NO
//写的时候读 NO


var WR_Wmutex *sync.RWMutex;
//var WR_Rmutex *sync.RWMutex;
var  WRch chan int

func main() {
	WR_Wmutex=new(sync.RWMutex)
	fmt.Println("start")
	WRch=make(chan  int )

	for i:=0;i<5;i++  {
		go TestRW_WMutex(i)
		go TestRW_RMutex(i)
	}
	for i:=0;i<10 ; i++ {
		<-WRch
	}
}

func TestRW_WMutex(index int) {
	WR_Wmutex.Lock();
	//写东西..
	fmt.Println()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"BBB 测试写，正在写东西.....","index=",index)
	time.Sleep(1*time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"BBB 测试写，即将离开写锁","index=",index)
	WR_Wmutex.Unlock()
	WRch<-1
}

func TestRW_RMutex(index int) {
	WR_Wmutex.RLock();
	//写东西..
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"AAA 测试读，正在读东西.....","index=",index)
	time.Sleep(1*time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),"AAA 测试读，即将离开读锁","index=",index)
	WR_Wmutex.RUnlock()
	WRch<-1
}
