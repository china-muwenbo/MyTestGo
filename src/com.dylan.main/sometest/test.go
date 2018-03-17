package main

import "fmt"

func main() {
	var m1 map[string] string
	m1 =make(map[string]string)
	m1["a"] ="aa"
	m1["b"] ="bb"

	m2 := make (map[string]string)
	m2["a"]="a"
	m2["b"] = "b"

	//m3:=map[string]string{
	//	"a":"aa",
	//	"b":"bb",
	//}
	if v, ok := m1["a"]; ok {
		fmt.Println(v)
	}else {
		fmt.Println("key not found")
	}

	for  k,v:=range m1 {
		fmt.Println(k,v)
	}
	for k,v:=range m1  {
		fmt.Println(k,v)
	}
	delete(m1,"a")
	delete(m1,"a")

}
