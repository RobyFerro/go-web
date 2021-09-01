package config

var RedisConf = struct {
	Host     string
	Port     int
	User     string
	Password string
}{
	Port: 6379,
}
