package config

var ElasticConf = struct {
	Hosts []string
}{
	Hosts: []string{
		"http://elastic:9200",
	},
}
