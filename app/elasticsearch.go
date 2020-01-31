package app

import (
	"github.com/elastic/go-elasticsearch/v8"
	"ikdev/go-web/app/config"
	"ikdev/go-web/exception"
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
