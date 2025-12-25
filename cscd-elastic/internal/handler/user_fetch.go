package handler

import (
	"context"

	"github.com/nats-io/nats.go"

	"github.com/kortelyov/cscd/cscd-elastic/internal/provider"
)

type UserFetchHandler struct {
	provider *provider.ElasticProvider
}

func NewUserFetchHandler(provider *provider.ElasticProvider) *UserFetchHandler {
	return &UserFetchHandler{
		provider: provider,
	}
}

func (h *UserFetchHandler) HandleUserFetch(msg *nats.Msg) {
	err := h.provider.FetchUser(context.Background(), msg.Subject)
	if err != nil {

	}
	msg.Respond(nil)
}
