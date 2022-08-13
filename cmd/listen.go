package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func WaitData(handler CommandHandler) *cobra.Command {
	if handler == nil {
		panic("internal error: handler can not be nil")
	}
	var listen = &cobra.Command{
		Use:   "listen",
		Short: "Listen data across nats system",
		Long: "Listen some topic from nats system received data will be write into write flag.\n" +
			"Core NATS offers an at most once quality of service.\n" +
			"If a subscriber is not listening on the subject (no subject match), \n" +
			"or is not active when the message is sent, the message is not received.",
	}
	topic := listen.Flags().StringP(topicFlag, string(topicFlag[0]), "", "topik name")
	host := listen.Flags().String(hostFlag, "127.0.0.1:4222", "127.0.0.1:4222")
	writer := listen.Flags().StringP(writeFlag, string(writeFlag[0]), "/dev/stdout", "/path/to/file")
	listen.Example = fmt.Sprintf("listen --%s my-topic", topicFlag)
	listen.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nats.Connect(*host, nats.Name(listen.Use), nats.Timeout(10*time.Second), nats.ReconnectWait(time.Second))
		if err != nil {
			return err
		}
		defer conn.Close()
		file, err := os.OpenFile(*writer, os.O_RDWR, 0)
		if err != nil {
			return err
		}
		defer file.Close()
		return handler(conn, *topic, file)
	}
	return listen
}
