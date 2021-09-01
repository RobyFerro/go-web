package config

import "os"

type SQLConf struct {
	Driver   string
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func GetSQL() *SQLConf {
	return &SQLConf{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     3306,
		Database: os.Getenv("MYSQL_DATABASE"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	}
}
