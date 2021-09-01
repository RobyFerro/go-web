package config

import "os"

var SQLConf = struct {
	Driver   string
	Host     string
	Port     int
	Database string
	User     string
	Password string
}{
	Driver:   "mysql",
	Host:     "localhost",
	Port:     3306,
	Database: os.Getenv("MONGO_DATABASE"),
	User:     os.Getenv("MYSQL_USER"),
	Password: os.Getenv("MYSQL_PASSWORD"),
}
