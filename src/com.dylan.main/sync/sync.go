package main

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

func Afunction(shownum int) {
	fmt.Println("开始",shownum)
	waitgroup.Done() //任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
	fmt.Println("结束",shownum)
}

func main() {

	for i := 0; i < 100; i++ {
		waitgroup.Add(1) //每创建一个goroutine，就把任务队列中任务的数量+1
		go Afunction(i)
		go Afunction(i)
		go Afunction(i)
	}
	waitgroup.Wait() //.Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}