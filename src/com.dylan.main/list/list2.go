package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	//把4元素放在最后
	e4 := l.PushBack(4)
	//把1元素放在最前
	e1:=l.PushFront(1)
	//在e4元素前面插入3
	l.InsertBefore(nil, e4)
	//在e4元素前面插入3
	l.InsertAfter("123132", e1)
	fmt.Println(" 元素 ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value," ")
	}
}
