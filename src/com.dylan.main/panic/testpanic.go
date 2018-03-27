package main


import (
	"fmt"
	//"errors"
	"errors"
)

func main() {
		err:=testError()
	    	if err !=nil {
	        fmt.Println(err)
	        }else{
	        	fmt.Println("程序没有错")
			}
	afterErrorfunc()
}

func testError() (error) {
	var err error
	defer  func ()   {
		if r := recover(); r != nil {
			fmt.Println("testError() 遇到错误:", r)
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("")
			}
			if err != nil {
				fmt.Println("recover后的错误:",err)
			}

		}
	}()
	panic(" \"panic 错误\"")
	fmt.Println("抛出一个错误后继续执行代码")
	return err
}


func afterErrorfunc(){
	fmt.Println("遇到错误之后 func ")
}