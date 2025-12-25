package handler

import (
	"context"
	"encoding/json"

	"github.com/kortelyov/cscd/cscd-contracts/pkg/contracts"
	"github.com/nats-io/nats.go"

	"github.com/kortelyov/cscd/cscd-elastic/internal/provider"
)

type UserPutHandler struct {
	provider *provider.ElasticProvider
}

func NewUserPutHandler(provider *provider.ElasticProvider) *UserPutHandler {
	return &UserPutHandler{
		provider: provider,
	}
}

func (h *UserPutHandler) HandleUserPut(msg *nats.Msg) {
	var user contracts.User
	_ = json.Unmarshal(msg.Data, &user)
	_ = h.provider.PutUser(context.Background(), &user)

	return
}
