package libs

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/iris-contrib/mail"
	"github.com/kataras/iris"
)

type Person struct {
	UserName string
}

func SendEmail(to []string, content string) {

	// config service
	cfg := mail.Config{
		Host:     cfg.Mail.Smtp,
		Username: cfg.Mail.User,
		Password: cfg.Mail.Password,
		Port:     cfg.Mail.Port,
	}

	// create the service
	mailService := mail.New(cfg)

	// declare to multiple destinations
	// var to = []string{"ivan.cduran@gmail.com"}

	// easy way
	//iris.Must(mailService.Send("iris e-mail test subject", "</h1>outside of context before server's listen!</h1>", to...))

	var tpl bytes.Buffer
	t := template.New("Simple email")
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: to[0]}
	t.Execute(&tpl, p)

	err := mailService.Send("Simple Email", tpl.String(), to...)

	if err != nil {
		fmt.Println("SendEmail: Fail")
	} else {
		fmt.Println("SendEmail: Success")
	}

}

func SendEmailComplex(to []string) {

	// config service
	cfg := mail.Config{
		FromAlias: cfg.Mail.FromAlias,
		Host:      cfg.Mail.Smtp,
		Username:  cfg.Mail.User,
		Password:  cfg.Mail.Password,
		Port:      cfg.Mail.Port,
	}

	mailService := mail.New(cfg)

	content := iris.TemplateString("emails/sample.html", iris.Map{
		"Message": " his is the rich message body sent by a template!!",
		"Footer":  "The footer of this e-mail!",
	}, iris.RenderOptions{"charset": "UTF-8"})

	err := mailService.Send("iris e-mail just t3st subject", content, to...)

	if err != nil {
		fmt.Println("SendEmail: Fail " + err.Error())
	} else {
		fmt.Println("SendEmail: Success")
	}

}
