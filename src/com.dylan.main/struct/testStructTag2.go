package main

import (
	"reflect"
	"fmt"
)

type User struct {
name string `name:"jerrytag"`
age int `age:"1"`
	}

func main() {
	user:=User{"jerry",20}
	field, ok := reflect.TypeOf(&user).Elem().FieldByName("name")
	if !ok {
		panic("Field not found")
	}
	fmt.Println(getStructTag(field))
}
func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
