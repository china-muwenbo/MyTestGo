package main

import "fmt"

type Base struct {
	basename string
}

//go语言可以让我们在声明struct时无需指定字段名，只要提供字段类型即可，这种就是匿名字段。
// 匿名字段的数据类型必须是一个具名类型或者指向具名类型的指针
type Derive1 struct {
	Base
}
type Derive2 struct {
	a *Base
}
type Derive3 struct {
	base Base
}

type Derive4 struct {
	base *Base
}

func main() {

	d:=new(Derive2)
	(*(d.a)).basename="hello"
	fmt.Println(d)

}

