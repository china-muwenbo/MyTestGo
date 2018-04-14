package main

import (
	"fmt"
	"gopkg.in/mgo.v2"

	"log"
	"os"
	"io"
)

func main() {
	//upload()
	ReadAll()
}
func downLoad ()  {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
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
	//通过文件名获取mp3
	file, err := session.DB("gridfs").GridFS("fs").Open("someThing.mp3")
	if err !=nil {
		fmt.Println(err)
		return
	}
	out, _ := os.OpenFile("E:\\test\\something.mp3.", os.O_CREATE, 0666)
	_,err = io.Copy(out, file)
	check(err)
	err = file.Close()
	check(err)
	err=out.Close()
	check(err)
	//b := make([]byte, 8192)
	//n, err := file.Read(b)
	//check(err)
	//fmt.Println(string(b))
	//check(err)
	//err = file.Close()
	//check(err)
	//fmt.Printf("%d bytes read\n", n)
}


func upload(){
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
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
	//通过文件名获取mp3
	file, err := session.DB("gridfs").GridFS("fs").Create("my.mp3")
	if err !=nil {
		fmt.Println(err)
		return
	}
	out, _ := os.OpenFile("E:\\test\\something.mp3.", os.O_RDWR, 0666)
	_,err = io.Copy(file, out)
	check(err)
	err = file.Close()
	check(err)
	err=out.Close()
	check(err)
}
func remove()  {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
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
	//通过文件名获取mp3
	err = session.DB("gridfs").GridFS("fs").Remove("my.mp3")
	if err !=nil {
		fmt.Println(err)
	}else {
		fmt.Print("刪除成功")
	}
}
//查询操作
func ReadAll()  {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
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
	//通过文件名获取mp3
	gfs := session.DB("gridfs").GridFS("fs");
	iter := gfs.Find(nil).Iter()

	result:=new(fileinfo)
	for iter.Next(&result) {
		fmt.Println("一个一个输出：", result)
	}

}
type fileinfo struct {
	//文件大小
	LENGTH int32
	//md5
	MD5 string
	//文件名
	FILENAME string
}

func check(err error){
	log.Print(err)
}
