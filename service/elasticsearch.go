package service

import (
	"github.com/RobyFerro/go-web/config"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/gommon/log"
)

// ConnectElastic returns an ElasticSearch client
func ConnectElastic() *elasticsearch.Client {
	elasticConf := config.GetElastic()
	cfg := elasticsearch.Config{
		Addresses: elasticConf.Hosts,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Error(err)
	}

	return es
}
