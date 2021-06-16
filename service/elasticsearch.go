package service

import (
	"github.com/RobyFerro/go-web/app"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/gommon/log"
)

func ConnectElastic(conf *app.Conf) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: conf.Elastic.Hosts,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Error(err)
	}

	return es
}
