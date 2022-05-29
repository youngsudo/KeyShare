package email

import (
	"fmt"

	"github.com/go-gomail/gomail"
)

func SendMailByGomail(from, to, password, emclient, header string, emport int, body *string) error {
	m := gomail.NewMessage()                 // 创建一个新的邮件
	m.SetAddressHeader("From", from, "auth") //发件人
	m.SetHeader("To", to)                    //收件人
	m.SetHeader("Subject", header)           //主题
	// m.Embed("./1.png")                       //附件,图片
	m.SetBody("text/html", fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<style>
		body{
			text-align: center;
		}
	</style>
	<body>
		<h1>欢迎来到共享密码(KeyShare)</h1>
		<p>%v</p>
	</body>
	</html>
	`, *body)) //正文

	/*
	   创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、密码，如果端口号为465的话，
	   自动开启SSL，这个时候需要指定TLSConfig
	*/
	d := gomail.NewDialer(emclient, emport, from, password) //创建发送邮件的对象

	if err := d.DialAndSend(m); err != nil { //发送邮件
		return err
	}
	return nil
}
