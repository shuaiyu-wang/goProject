package main

import (
	"fmt"
	"goProject/email"
)

func main() {
	//定义收件人
	mailTo := []string{
		"1318008936@qq.com",
	}
	//邮件主题为"Hello"
	subject := "BUG"
	// 邮件正文
	body := "Hello,zmf"

	err := email.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println("send fail", err)
		return
	}
	fmt.Println("send successfully")

}
