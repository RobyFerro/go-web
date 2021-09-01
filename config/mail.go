package config

import "os"

type MailConf struct {
	From     string
	Host     string
	Username string
	Password string
	Port     int
}

func GetMail() *MailConf {
	return &MailConf{
		From:     "pirate@goweb.dev",
		Host:     "localhost",
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Port:     25,
	}
}
