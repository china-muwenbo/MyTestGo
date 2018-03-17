package main

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

func main() {
	//接受函数 sayHello
	http.HandleFunc("/", sayHello)
	err:=http.ListenAndServe(":9090",nil)
	if err !=nil{
		log.Fatal("ListenServe ", err )
	}
}

func sayHello(resp http.ResponseWriter ,req *http.Request){
	req.ParseForm()
	fmt.Println("path",req.URL.Path)
	fmt.Println("scheme",req.URL.Scheme)
	fmt.Println("opaque",req.URL.Opaque)
	fmt.Println(req.Form["url_long"])
	for  k,v:=range req.Form  {
		fmt.Println("key",k)
		fmt.Println("val:",strings.Join(v," "))
	}
	fmt.Fprintf(resp,"hello  dylan ")
}

func test(a int, i int) {
	fmt.Println(i)
}
