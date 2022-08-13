package main

import "github.com/blabu/nts/cmd"

func main() {
	root := cmd.CreateRootCmd()
	root.AddCommand(cmd.SendData(sendFileHandler))
	root.AddCommand(cmd.WaitData(listenTopicHandler))
	root.AddCommand(cmd.StreamData(streamHandler))
	root.Execute()
}
