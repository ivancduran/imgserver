package libs

import (
	"os"

	"github.com/joho/godotenv"
)

type Conf struct {
	Devel  string
	Domain string
	Host   string
	Api    string
	Fs     string
	Db     Db
	Mail   MailConf
}

type Db struct {
	Host         string
	Database     string
	AuthDatabase string
	AuthUsername string
	AuthPassword string
}

type MailConf struct {
	Smtp      string
	FromAlias string
	User      string
	Password  string
	Port      string
}

func GetConf() Conf {
	godotenv.Load("env", "env.prod")

	gC := Conf{}

	gC.Devel = os.Getenv("Devel")
	gC.Domain = os.Getenv("Domain")
	gC.Host = os.Getenv("Host")
	gC.Api = os.Getenv("Api")
	gC.Fs = os.Getenv("Fs")
	gC.Db.Host = os.Getenv("DB_Host")
	gC.Db.Database = os.Getenv("DB_Database")
	gC.Db.AuthDatabase = os.Getenv("DB_AuthDb")
	gC.Db.AuthUsername = os.Getenv("DB_User")
	gC.Db.AuthPassword = os.Getenv("DB_Pass")
	gC.Mail.Smtp = os.Getenv("Mail_Server")
	gC.Mail.FromAlias = os.Getenv("Mail_From")
	gC.Mail.User = os.Getenv("Mail_User")
	gC.Mail.Password = os.Getenv("Mail_Pass")
	gC.Mail.Port = os.Getenv("Mail_Port")

	// fmt.Println(os.Getenv("DB_Host"))
	// fmt.Println(os.Getenv("DB_Database"))
	// fmt.Println(os.Getenv("DB_AuthDb"))
	// fmt.Println(os.Getenv("DB_User"))
	// fmt.Println(os.Getenv("DB_Pass"))

	return gC
}

var cfg = GetConf()
