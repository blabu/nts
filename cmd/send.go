package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func SendData(handler CommandHandler) *cobra.Command {
	if handler == nil {
		panic("internal error: handler can not be nil")
	}
	var send = &cobra.Command{
		Use:   "send",
		Short: "Send file across nats system over steaming",
		Long: "Send some file to the topic. \n" +
			"If file is being larger than 1 Mb (default max payload size) it will be sliced.\n" +
			"Core NATS offers an at most once quality of service. \n" +
			"If a subscriber is not listening on the subject (no subject match), \n" +
			"or is not active when the message is sent, the message is not received.\n",
	}
	topic := send.Flags().StringP(topicFlag, string(topicFlag[0]), "", "topik name")
	host := send.Flags().String(hostFlag, "127.0.0.1:4222", "127.0.0.1:4222")
	reader := send.Flags().StringP(readFlag, string(readFlag[0]), "/dev/stdin", "/path/to/file")
	send.Example = fmt.Sprintf("sendf --%s my-topic --%s /path/to/file", topicFlag, readFlag)
	send.RunE = func(cmd *cobra.Command, args []string) error {
		conn, err := nats.Connect(*host, nats.Name(send.Use), nats.Timeout(10*time.Second), nats.ReconnectWait(time.Second))
		if err != nil {
			return err
		}
		defer conn.Close()
		file, err := os.OpenFile(*reader, os.O_RDONLY, 0)
		if err != nil {
			return err
		}
		defer file.Close()
		return handler(conn, *topic, file)
	}
	return send
}
