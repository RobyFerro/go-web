package config

type MongoConf struct {
	Database string
	Host     string
	Port     int
	User     string
	Password string
}

func GetMongo() *MongoConf {
	return &MongoConf{
		Database: "go_queue",
		Host:     "mongo",
		Port:     27017,
		User:     "",
		Password: "",
	}
}
