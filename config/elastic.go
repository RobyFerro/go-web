package config

type ElasticConf struct {
	Hosts []string
}

func GetElastic() *ElasticConf {
	return &ElasticConf{
		Hosts: []string{
			"http://elastic:9200",
		},
	}
}
