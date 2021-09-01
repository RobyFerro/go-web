package config

type redisConf struct {
	Host     string
	Port     int
	User     string
	Password string
}

func GetRedis() *redisConf {
	return &redisConf{
		Port: 6379,
	}
}
