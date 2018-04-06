package main

import (
	fmt "fmt"
)
type HmConn struct {

}
type TZ int


func main() {
	var z TZ
	z.Print()
	(*TZ).Print(&z)
}
func(a *TZ)Print(){
	fmt.Println("TZ")
}


