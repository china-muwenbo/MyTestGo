package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func main() {
	session, err := mgo.Dial("mongodb://123.207.215.200:27017")
	defer session.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	names, err := session.DatabaseNames();
	if err != nil {
		fmt.Println("未查询到数据库名字:", err)
	}
	fmt.Println(names)
	session.DB("mytest").C("people").Insert()
}



