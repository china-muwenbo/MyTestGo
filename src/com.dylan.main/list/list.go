package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 创建一个 list
	l := list.New()
	//把4元素放在最后
	e4 := l.PushBack(4)
	//把1元素放在最前
	e1 := l.PushFront(1)
	//在e4元素前面插入3
	l.InsertBefore(3, e4)
	//在e1后面插入2
	e2:=l.InsertAfter(2, e1)
	// 遍历所有元素并打印其内容
	fmt.Println(" 元素 ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	//获取l 最前的元素
	et1:=l.Front();
	fmt.Println("list 最前的元素 Front  ",et1.Value)
	//获取l 最后的元素
	et2:=l.Back()
	fmt.Println("list 最后的元素  Back ",et2.Value)
	//获取l的长度
	fmt.Println("list 的长度为： Len ",l.Len())
	//向后移动
	l.MoveAfter(e1,e2)
	fmt.Println("把1元素移动到2元素的后面 向后移动后 MoveAfter :")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	//向前移动
	l.MoveBefore(e1,e2)
	fmt.Println("\n把1元素移动到2元素的前面 向前移动后 MoveBefore :")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	//移动到最后面
	l. MoveToBack(e1)
	fmt.Println("\n 1元素出现在最后面 MoveToBack ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	//移动到最前面
	l. MoveToFront(e1)
	fmt.Println("\n 1元素出现在最前面 MoveToFront ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	//删除元素
	fmt.Println("")
	l.Remove(e1)
	fmt.Println("\n e1元素移除后 Remove ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	// init 可以用作 clear
	l.Init()
	fmt.Println("\n list init()后 ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
	fmt.Println("list 的长度Init  ",l.Len())
	//向前移动
	//

}