package cmd

import (
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func StreamData(handler CommandHandler) *cobra.Command {
	if handler == nil {
		panic("internal error: handler can not be nil")
	}
	stream := &cobra.Command{
		Use:   "stream",
		Short: "Open a jet stream, see the documentation https://docs.nats.io/nats-concepts/jetstream",
		Long: "If you need higher qualities of service (at least once and exactly once), \n" +
			"or functionalities such as persistent streaming, de-coupled flow control, and Key/Value Store, you can use NATS JetStream,\n" +
			"which is built in to the NATS server (but needs to be enabled).",
	}
	topic := stream.Flags().StringP(topicFlag, string(topicFlag[0]), "", "topik name")
	host := stream.Flags().String(hostFlag, "127.0.0.1:4222", "127.0.0.1:4222")
	stream.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nats.Connect(*host, nats.Name(stream.Use), nats.Timeout(10*time.Second), nats.ReconnectWait(time.Second))
		if err != nil {
			return err
		}
		defer conn.Close()
		return handler(conn, *topic, nil)
	}
	return stream
}
