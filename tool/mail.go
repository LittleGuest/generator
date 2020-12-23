package tool

import (
	"gopkg.in/gomail.v2"
)

// SendTextEmail send text email
func SendTextEmail(to string, subject string, content string) error {
	mmsg := gomail.NewMessage()
	mmsg.SetHeader("From", "2190975784@qq.com")
	mmsg.SetHeader("To", to)
	mmsg.SetHeader("Subject", subject)
	mmsg.SetBody("text/plain", content)

	gd := gomail.NewDialer("smtp.qq.com", 25, "2190975784@qq.com", "enomoilgbvckdice")
	return gd.DialAndSend(mmsg)
}
