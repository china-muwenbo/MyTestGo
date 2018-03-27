package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	name string `name:"jerry"`
	age int `age:"28" `
}

type T struct {
	A int `A:"1"`
	B string `B:"a"`
}

func main() {

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}


}

