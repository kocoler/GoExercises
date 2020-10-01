package main

import (
	"bytes"
	"fmt"
	"github.com/go-gomail/gomail"
	"log"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {

	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",
	//  "port": "465",
	//}

	mailConn := map[string]string{
		"user": "tsglsdrs@163.com",
		"pass": "ccnu2020dx",
		"host": "smtp.163.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "孤独星球"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
func main() {

	mailTo := []string{
		"tsglsdrs@163.com",
	}

	subject := "孤独星球　新消息提醒"

	url := "http://localhost:3000/forgot?method=email&id=20&code=PERWYqZzWM6zaadOnxgx&username=wwwwww"
	var bt bytes.Buffer
	bt.WriteString("<a href=\"")
	bt.WriteString(url)
	bt.WriteString("\"")
	bt.WriteString(">")
	bt.WriteString(url)
	bt.WriteString("</a>")
	//body := "<a href=\""+"https://www.baidu.com/"+"\">"+"oo"+"</a>"

	err := SendMail(mailTo, subject, bt.String())
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}
