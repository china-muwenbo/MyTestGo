package main

import "fmt"

type person struct {
	name string
	ctiy string
}




func main() {
	var p *person
	p = &person{
	}
	(*p).name = "Jerry"
	p.name="yy"
	(*p).ctiy = "xi'an"
	fmt.Println(p)
	//Pp(p)
	p.Pp()
	fmt.Println(p)
}

func(p *person) Pp() {
	p.name="func"
	fmt.Println("func p ", p)
}

//
//package main
//
//import "fmt"
//
//type person struct {
//	name string
//	ctiy string
//}
//
//func main() {
//	p := &person{
//	}
//	p.name = "Jerry"
//	p.ctiy = "xi'an"
//	fmt.Println(p)
//	Pp(p)
//	fmt.Println(p)
//}
//func Pp(p *person) {
//	p.name="func"
//	fmt.Println("func p ", p)
//}