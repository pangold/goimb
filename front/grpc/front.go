package grpc

import (
	"context"
	"fmt"
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"google.golang.org/grpc"
	"log"
)

type FrontApiAdapter struct {
	conf config.Host
	conn *grpc.ClientConn
	context context.Context
	api protocol.ImApiServiceClient
}

func NewFrontApi(conf config.Host) *FrontApiAdapter {
	conn, err := grpc.Dial(conf.Address, grpc.WithInsecure())
	if err != nil {
		panic("connect to grpc api server error: " + err.Error())
	}
	front := &FrontApiAdapter{
		conf: conf,
		conn: conn,
		context: context.Background(),
		api: protocol.NewImApiServiceClient(conn),
	}
	return front
}

func (this *FrontApiAdapter) Send(message *protocol.Message) error {
	stream, err := this.api.Send(this.context)
	if err != nil {
		return fmt.Errorf("client send error: %v", err)
	}
	if err = stream.Send(message); err != nil {
		return fmt.Errorf("send error: %v", err)
	}
	return nil
}

func (this *FrontApiAdapter) SendEx(messages []*protocol.Message) error {
	// temp implement
	for _, message := range messages {
		return this.Send(message)
	}
	return nil
}

func (this *FrontApiAdapter) Broadcast(message *protocol.Message) error {
	stream, err := this.api.Broadcast(this.context)
	if err != nil {
		return fmt.Errorf("client broadcast error: %v", err)
	}
	if err = stream.Send(message); err != nil {
		return fmt.Errorf("broadcast error: %v", err)
	}
	return nil
}

func (this *FrontApiAdapter) BroadcastEx(messages []*protocol.Message) error {
	// temp implement
	for _, message := range messages {
		return this.Broadcast(message)
	}
	return nil
}

func (this *FrontApiAdapter) GetConnections() []string {
	res, err := this.api.GetConnections(this.context, &protocol.Empty{})
	if err != nil {
		log.Printf("failed to get connections, error: %v", err)
		return nil
	}
	return res.GetUserIds()
}

func (this *FrontApiAdapter) Online(uid string) bool {
	res, err := this.api.Online(this.context, &protocol.User{UserId: uid})
	if err != nil {
		log.Printf("failed to get online users(%s), error: %v", uid, err)
		return false
	}
	return res.GetSuccess()
}

func (this *FrontApiAdapter) Kick(uid string) bool {
	res, err := this.api.Kick(this.context, &protocol.User{UserId: uid})
	if err != nil {
		log.Printf("failed to kick user(%s), error: %v", uid, err)
		return false
	}
	return res.GetSuccess()
}

