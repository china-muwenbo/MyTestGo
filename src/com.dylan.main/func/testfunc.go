package main

import "fmt"

// 学习匿名函数和闭包
func main() {
	f:=getfunc()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(i)
}
//闭包用一个专业一点的说法就是：函数调用返回后一个没有释放资源的栈区。
var i int
func getfunc() func() int{
	i=-1
	return func() int {
		i=i+1
		return i
	}
}
//  运行结果1 ，2，3

