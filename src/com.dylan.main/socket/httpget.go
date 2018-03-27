package main
import (
	"net/http"
	"fmt"
	"io/ioutil"
)
func main()  {
	get()
}
func get(){
	resp,err :=    http.Get("http://www.baidu.com")
	if err!=nil {
		//如果没有获取到url数据 会返回一个error 类型错误
		fmt.Println("error=",err,";")
	}
	defer resp.Body.Close()
	//提取响应头数据
	b,err:=ioutil.ReadAll(resp.Body)
	//将字节切片转成成String 输出
	fmt.Print(string(b))
}