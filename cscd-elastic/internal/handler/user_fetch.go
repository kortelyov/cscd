package handler

import (
	"github.com/nats-io/nats.go"
)

type UserFetchHandler struct{}

func (h *UserFetchHandler) HandleUserGet(msg *nats.Msg) {
	return
}
