package main

import (
	"fmt"
	"math/rand"
	"time"
)
// go 协程应用之一
// 创建10个线程
//主线程阻塞，等待所有子线程计算完成
// 子线程完成把结果放在子线程中

func main()  {
	var testchan =make(chan int ,10)
	for i:=0;i<10 ;i++ {
		go child(testchan,i)
	}
	for i:=0;i<10 ;i++  {
		a:=<-testchan
		fmt.Println("协程完成:",a)
	}
	fmt.Println("主程完成")

}

func child(mychan chan int,index int  ){
	/*
	这里遇到一个问题 线程休眠需要一个 time.Duration类型的参数，也就是int64
	参数 此时必须将 int 转换成 time.Duration 参数类型
	用 time.Duration（int(要转换的数据)） 即可转换
	*/
	time.Sleep(time.Second*time.Duration(rand.Intn(10)))
	mychan<-index;

}
