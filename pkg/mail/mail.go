// Package mail 发送短信
package mail

import (
	"server/pkg/config"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte // Plaintext message (optional)
	HTML    []byte // Html message (optional)
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

// NewMailer 单例模式获取
func NewMailer() *Mailer {
	once.Do(func() {
		switch config.Get("mail.default") {
		case "smtp":
			internalMailer = &Mailer{
				Driver: &SMTP{},
			}
		case "ses":
			internalMailer = &Mailer{
				Driver: &SES{},
			}
		}
	})

	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.mailers."+config.Get("mail.default")))
}
