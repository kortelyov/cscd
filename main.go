package main

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/kortelyov/cscd/cscd-contracts/pkg/contracts"
	"github.com/kortelyov/cscd/cscd-contracts/pkg/subjs"
	"github.com/kortelyov/cscd/cscd-contracts/pkg/wrpr"
)

func main() {
	ctx := context.Background()

	nc, err := nats.Connect(nats.DefaultURL,
		nats.Timeout(10*time.Second),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			log.Printf("Disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			log.Print("Reconnected to NATS")
		}),
		nats.RetryOnFailedConnect(true),
	)

	user := contracts.User{
		Id:    "",
		Email: "",
		Roles: nil,
	}

	// 1. Прямой вызов дата-сервиса (Прагматично!)
	err = wrpr.Request(ctx, nc, subjs.SubjectElasticUserFetch,
		&contracts.GetUserRequest{Id: 0}, &user)

	if err != nil {
		// 2. Если юзера нет, создаем (логика прямо здесь)
		err = wrpr.Request(ctx, nc, subjs.SubjectElasticUserPut,
			&contracts.CreateUserRequest{User: &contracts.User{Id: "0"}}, "&protobuf.Empty{}")
	}
}
