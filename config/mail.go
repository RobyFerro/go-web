package config

import "os"

var MailConf = struct {
	From     string
	Host     string
	Username string
	Password string
	Port     int
}{
	From:     "pirate@goweb.dev",
	Host:     "localhost",
	Username: os.Getenv("MAIL_USERNAME"),
	Password: os.Getenv("MYSQL_PASSWORD"),
	Port:     25,
}
