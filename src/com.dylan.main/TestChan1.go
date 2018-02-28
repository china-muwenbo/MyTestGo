package main

import (
	"time"
	"fmt"
	"math/rand"
)

/*
chan 应用生产者和消费者模型
生产者每生产一个东西和消费者就消费一个东西
消费者消费东西速度过快，生产者者跟不上，消费者就阻塞.
生产者生产东西过快，消费者跟不上，生产者阻塞
叫做同步的生产与消费

如果想拥有一个生产缓冲区，缓存一下 可以设置chan 缓存 cpChan:=make(chan int ,4)
这样就变成非同步的生产与消费

*/

func main() {
	cpChan:=make(chan int)
	go  consumer(cpChan)
	go  producter(cpChan)

	time.Sleep(time.Second*10)
}

func producter(cpChan chan int) {
	for i:=0; i<10;i++  {
		fmt.Println("生产：",i)
		cpChan<-i
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(1000)))
	}
}

func consumer(cpChan chan int) {
	for i:=0; i<10;i++  {
		p1:=<-cpChan
		fmt.Println("消费了",p1)
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(1000)))
	}
}

