package cmd

import (
	"io"

	"github.com/nats-io/nats.go"
)

type CommandHandler func(con *nats.Conn, topic string, r io.ReadWriter) error
