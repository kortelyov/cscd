package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"

	"github.com/kortelyov/cscd/cscd-contracts/pkg/contracts"
)

type ElasticProvider struct {
	client *elasticsearch.Client
}

func NewElasticProvider(client *elasticsearch.Client) *ElasticProvider {
	return &ElasticProvider{
		client: client,
	}
}

func (p *ElasticProvider) FetchUser(ctx context.Context, username string) (*contracts.User, error) {
	req := esapi.XPackSecurityGetUserRequest{
		Username: []string{username},
	}

	response, err := req.Do(ctx, p.client)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var rspmap map[string]interface{}
	if err = json.Unmarshal(body, &rspmap); err != nil {
		return nil, err
	}

	user, err := convert(username, rspmap)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *ElasticProvider) PutUser(ctx context.Context, user *contracts.User) error {
	b, _ := json.Marshal(user)

	req := esapi.XPackSecurityPutUserRequest{
		Body:     bytes.NewReader(b),
		Username: user.Email,
	}

	_, err := req.Do(ctx, p.client)
	if err != nil {
		return err
	}

	return nil
}

func convert(username string, from map[string]interface{}) (*contracts.User, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}
	m, ok := from[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	var user contracts.User
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
