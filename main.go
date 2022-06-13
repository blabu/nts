package main

import "github.com/blabu/nats_test/cmd"

func main() {
	root := cmd.CreateRootCmd()
	root.AddCommand(cmd.SendDataCommand(sendFileHandler))
	root.AddCommand(cmd.WaitDataCommand(listenTopicHandler))
	root.Execute()
}
