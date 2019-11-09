package main

import (
	"context"
	"log"

	gapi "github.com/eviot/grpc-api-go"
)

func main() {
	plugin, err := gapi.NewPluginService(5551)
	if err != nil {
		log.Fatalln(err)
	}
	plugin.RegisterInputPipe("echo-pipe", &Echo{})
	plugin.WaitForClose()
}

type Echo struct {
}

func (p *Echo) ReceiveMsg(msg *gapi.ReceiveMsgReq) {
	log.Println(msg)
	gapi.SendMsg(context.Background(), msg.Pipe, "out", msg.Message)
}
