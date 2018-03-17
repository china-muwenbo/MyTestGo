package main

import (
	"gopkg.in/gomail.v2"
	"fmt"
)

func main() {
	m := gomail.NewMessage()

	m.SetHeader("From", "18729243053@163.com")

	m.SetHeader("To", "1061792506@qq.com")

	m.SetAddressHeader("Cc", "1679496833@qq.com", "Dan")
	//使用Subject 指明
	m.SetHeader("Subject", "Hello!")

	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// QQ邮箱 居然给老子标记垃圾邮件
	d := gomail.NewDialer("smtp.163.com", 465, "18729243053@163.com", "********")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}else {
		fmt.Println("发送完成")

	}
}