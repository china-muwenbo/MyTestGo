package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
type Test struct {
	name string
	firstName string
}


func main() {

	c, err := 	redis.Dial("tcp", "123.207.215.205:6379")
	 if err != nil {
		     fmt.Println(err)
		     return
		 }
	//t:=Test{"hello","123"}
	//t1:=&Test{"helo","123"}

	err=c.Send("auth","muwenbo")
	if err !=nil {
    fmt.Println(err)
    }
	err=c.Send("set","key","value")
	re,err:=redis.String(c.Do("get","key"))
	fmt.Println(re)
	defer c.Close()
}
