package config

var MongoConf = struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}{
	Database: "go_queue",
	Host:     "mongo",
	Port:     27017,
	User:     "",
	Password: "",
}
