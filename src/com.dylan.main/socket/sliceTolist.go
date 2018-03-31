package main

import (
	"time"
	"fmt"
	"container/list"
)

func main() {

	//t := time.Now()
	sli:=make([]int ,10)
	for i := 0; i<1*100000*1000;i++  {
	sli=append(sli, 1)
	}
	//fmt.Println("slice 创建速度：" + time.Now().Sub(t).String())

	//t = time.Now()
	l:=list.New()
	for i := 0; i<1*100000*1000;i++  {
		l.PushBack(1)
	}
	//fmt.Println("list 创建速度: " + time.Now().Sub(t).String())
	// 比较遍历

	t := time.Now()
	for _,_ = range sli {
		//fmt.Printf("values[%d]=%d\n", i, item)
	}
	fmt.Println("遍历slice的速度:" + time.Now().Sub(t).String())

	t = time.Now()
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
	}
	fmt.Println("遍历list的速度:" + time.Now().Sub(t).String())

	//比较插入
	t = time.Now()
	slif:=sli[:100000*500]
	slib:=sli[100000*500:]
	slif=append(slif, 10)
	slif=append(slif, slib...)
	//fmt.Println("slice 的插入速度" + time.Now().Sub(t).String())
	fmt.Println("list 的插入速度 : " ,time.Now().Sub(t).Nanoseconds())

	var em *list.Element
	len:=l.Len()
	var i int

	for e := l.Front(); e != nil; e = e.Next() {
		i++
		if i ==len/2 {
			em=e
			break
		}
	}
	//忽略掉找一个元素的速度。
	t = time.Now()
	ef:=l.PushBack(2)
	l.MoveBefore(ef,em)
	//l.insert()
	fmt.Println("list 的插入速度 : " ,time.Now().Sub(t).Nanoseconds())



}
