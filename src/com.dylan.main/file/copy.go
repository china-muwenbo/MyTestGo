package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
	"log"
	//"golang.org/x/text"
	"io"
)

func main() {

	var xnt string
	var new  string
	var video  string
	args:=os.Args
	if len(args)<3{
		fmt.Println("请输入合理参数")
		//old =args[1]
		//new =args[2]
		//return
	}

	fmt.Println("当前路径：",getCurrentDirectory())
	//old =getCurrentDirectory()+"\\"+args[1];
	fmt.Println(xnt)
	//new =getCurrentDirectory()+"\\"+args[2];
	fmt.Println("加密后的头文件的路径：",xnt)
	fmt.Println("视频文件路径：",video)
	var f *os.File
	var fvideo *os.File
	var fnew *os.File
	var err1 error

	xnt="E:\\goCode\\src\\MyTestGo\\src\\com.dylan.main\\file\\hello.xnt"
	video="E:\\goCode\\src\\MyTestGo\\src\\com.dylan.main\\file\\疯狂动物城15分钟.mkv"
	new="E:\\goCode\\src\\MyTestGo\\src\\com.dylan.main\\file\\疯狂动物城15分钟.bwv"

	if checkFileIsExist(xnt) { //如果文件存在
		f,err1=os.OpenFile(xnt, os.O_RDWR , 0666) //打开文件
		if err1 !=nil {
			fmt.Println("打开文件失败")
			return
			}
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
		return
	}
	buffer:=make([]byte,1024)
	n,err:=f.Read(buffer)
	if err !=nil {
	   fmt.Println("读取失败")
	   return
	}
	if n<1024{
		fmt.Println("头文件不应该小于1024字节")
		return
	}
	f.Close();

	if checkFileIsExist(video) { //如果文件存在
		fvideo,err1=os.OpenFile(video, os.O_RDWR , 0666) //打开文件
		if err1 !=nil {
			fmt.Printf("打开%s文件失败\n",video)
			return
		}
	} else {
		fmt.Println("视频文件不存在")
		return
	}

	if checkFileIsExist(new) { //如果文件存在
		fnew,err1=os.OpenFile(new, os.O_RDWR|os.O_CREATE , 0666) //打开文件
		if err1 !=nil {
			fmt.Printf("打开%s文件失败\n",new)
			return
		}
	} else {
		fnew,err1=os.Create(new);
		fmt.Println("创建新文件")
		//return
	}
	fnew.Write(buffer);
	copyBuffer:=make([]byte,1024*1024)
	fvideo.Seek(1024,0)
	for {
		length, err := fvideo.Read(copyBuffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("复制文件出错", err)
				return
			}
		}
		fnew.Write(copyBuffer[0:length])
		//fnew.Write(copyBuffer)
	}

	fnew.Close()

	fvideo.Close()


}
type A struct {
	Name string
	B string
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}



func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir //strings.Replace(dir, "\\", "/", -1)
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}


/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}