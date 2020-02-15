package service

import (
	"github.com/RobyFerro/go-web-framework"
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElastic(conf *gwf.Conf) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: conf.Elastic.Hosts,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		gwf.ProcessError(err)
	}

	return es
}
