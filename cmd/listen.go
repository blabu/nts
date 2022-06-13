package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

func WaitDataCommand(handler func(conn *nats.Conn, topic string, w io.Writer) error) *cobra.Command {
	var listen = &cobra.Command{
		Use:   "listen",
		Short: "listen data across nats system over steaming",
		Long:  "listen some topic from nats system",
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
