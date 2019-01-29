/*
@Time : 2019/1/23 下午3:32 
@Author : zwcui
@Software: GoLand
*/
package controllers

import (
	"net/smtp"
	"fmt"
	"strings"
	"anonymousFriends/util"
)




func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	auth := smtp.PlainAuth(
		"",
		"service@mail.zwcui.cn",
		"",
		"smtpdm.aliyun.com",
	)
	too := []string{"747660511@qq.com"}
	subject = "This is the email subject"
	message := "This is the email body."
	msg := fmt.Sprintf("To: %s\r\nFrom: service@mail.zwcui.cn\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", strings.Join(too, ","), subject, message)
	err := smtp.SendMail(
		"smtpdm.aliyun.com:25",
		auth,
		"service@mail.zwcui.cn",
		too,
		[]byte(msg),
	)
	if err != nil {
		util.Logger.Info(err.Error())
	}

	return err
}












