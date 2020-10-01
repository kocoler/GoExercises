package service

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/go-gomail/gomail"
	"strconv"
)

func SendResetPasswordMail(email, memberId, url string) error {

	mailConn := map[string]string{
		"user": "no-reply@casbin.com",
		"pass": "1qaz2wsx.",
		"host": "smtp.qiye.aliyun.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	mail := gomail.NewMessage()

	var bt bytes.Buffer

	bt.WriteString("Hi: ")
	bt.WriteString(memberId); bt.WriteString(","); bt.WriteString("<br/><br/>")
	bt.WriteString("我们的系统收到一个请求，说你希望通过电子邮件重新设置你在 ")
	name := beego.AppConfig.String("appname")
	bt.WriteString(name)
	bt.WriteString(" 的密码。你可以点击下面的链接开始重设密码："); bt.WriteString("<br/><br/>")
	bt.WriteString("<a href=\"")
	bt.WriteString(url)
	bt.WriteString("\">")
	bt.WriteString(url)
	bt.WriteString("</a>")
	bt.WriteString(url); bt.WriteString("<br/><br/>")
	bt.WriteString("如果这个请求不是由你发起的，那没问题，你不用担心，你可以安全地忽略这封邮件。 "); bt.WriteString("<br/><br/>")
	bt.WriteString("如果你有任何疑问，可以回复这封邮件向我们提问。 "); bt.WriteString("<br/><br/>")
	bt.WriteString("<front color=\"#888888\">")
	bt.WriteString(name)
	bt.WriteString("</font>")
	body := bt.String()
	//fmt.Println(body)
	mail.SetHeader("From", mail.FormatAddress(mailConn["user"], name))
	mail.SetHeader("To", email)    //发送给多个用户
	mail.SetHeader("Subject", "["+name+"]"+" 重设密码") //设置邮件主题
	mail.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(mail)
	return err

}
