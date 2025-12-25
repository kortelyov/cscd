package wrpr

import (
	"context"

	"github.com/nats-io/nats.go"

	"google.golang.org/protobuf/proto"
)

func Request[T any, R any](ctx context.Context, nc *nats.Conn, subj string, req T, res R) error {
	data, err := proto.Marshal(any(req).(proto.Message))
	if err != nil {
		return err
	}

	msg, err := nc.RequestWithContext(ctx, subj, data)
	if err != nil {
		return err
	}

	return proto.Unmarshal(msg.Data, any(res).(proto.Message))
}
