package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/kortelyov/cscd/cscd-contracts/pkg/subjs"
	"github.com/kortelyov/cscd/cscd-elastic/internal/handler"
)

func main() {
	logger := log.New(os.Stdout, "[CSCD-ELASTIC] ", log.LstdFlags)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	nc, err := nats.Connect(os.Getenv("NATS_URL"),
		nats.Name("cscd-elastic"),
		nats.Timeout(10*time.Second),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			log.Printf("disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			logger.Print("reconnected to NATS")
		}),
		nats.RetryOnFailedConnect(true),
	)
	if err != nil {
		logger.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	fetch, err := nc.QueueSubscribe(subjs.SubjectElasticUserFetch, subjs.QueueElastic, handler.HandleUserFetch)
	if err != nil {
		logger.Fatalf("subscription error: %v", err)
	}
	put, err := nc.QueueSubscribe(subjs.SubjectElasticUserPut, subjs.QueueElastic, handler.HandleUserPut)
	if err != nil {
		logger.Fatalf("subscription error: %v", err)
	}

	<-ctx.Done()
	logger.Print("shutdown signal received, commencing graceful shutdown...")

	if err = fetch.Drain(); err != nil {
		logger.Printf("error during sub drain: %v", err)
	}
	if err = put.Drain(); err != nil {
		logger.Printf("error during sub drain: %v", err)
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan struct{})

	go func() {
		close(done)
	}()

	select {
	case <-done:
		logger.Print("all resources closed successfully")
	case <-shutdownCtx.Done():
		logger.Printf("shutdown timed out: %v, forcing exit", shutdownCtx.Err())
	}
}
