package main

import (
	"reflect"
	"fmt"
	"strings"
	"log"
)

func main() {
	fmt.Println(GetFieldName(Student{}))
	fmt.Println(GetFieldName(&Student{}))
	fmt.Println(GetFieldName(""))

	fmt.Println(GetTagName(&Student{}))
}


type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}



//获取结构体中字段的名称
func GetFieldName(structName interface{}) []string  {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i:= 0;i<fieldNum;i++ {
		result = append(result,t.Field(i).Name)
	}
	return result
}


//获取结构体中Tag的值，如果没有tag则返回字段值
func GetTagName(structName interface{}) []string  {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string,0,fieldNum)
	for i:= 0;i<fieldNum;i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			tagName = tags[1]
		}
		result = append(result,tagName)
	}
	return result
}
