package app

import (
	"github.com/RobyFerro/go-web/app/config"
	"github.com/RobyFerro/go-web/exception"
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElastic(conf config.Conf) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: conf.Elastic.Hosts,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		exception.ProcessError(err)
	}

	return es
}
