package handler

import (
	"context"

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
	err := h.provider.PutUser(context.Background(), msg.Subject)

	return
}
