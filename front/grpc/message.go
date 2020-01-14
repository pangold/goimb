package grpc

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type MessageAdapter struct {
	conn *grpc.ClientConn
	context context.Context
	dispatcher protocol.ImDispatchServiceClient
}

func NewMessageAdapter(conf config.Host) *MessageAdapter {
	conn, err := grpc.Dial(conf.Address, grpc.WithInsecure())
	if err != nil {
		panic("connect to message dispatch grpc server error: " + err.Error())
	}
	ctx := context.Background()
	return &MessageAdapter{
		conn: conn,
		context: ctx,
		dispatcher: protocol.NewImDispatchServiceClient(conn),
	}
}

func (this *MessageAdapter) SetMessageHandler(handler func(*protocol.Message) error) {
	go this.on_message(handler)
}

// handle dispatched messages
func (this *MessageAdapter) on_message(cb func(*protocol.Message) error) {
	cli, _ := this.dispatcher.Dispatch(context.Background(), &protocol.Empty{})
	for {
		message, err := cli.Recv()
		if err != nil {
			log.Printf("get dispatched message error: %v", err)
			break
		}
		cb(message)
	}
}