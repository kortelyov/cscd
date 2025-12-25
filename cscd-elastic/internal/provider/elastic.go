package provider

import (
	"context"

	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type ElasticProvider struct {
	client *elastic.Client
}

func (p *ElasticProvider) FetchUser(ctx context.Context, id string) error {

	return nil
}
