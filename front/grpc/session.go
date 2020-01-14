package grpc

import (
	"gitlab.com/pangold/goim/protocol"
	"gitlab.com/pangold/goimb/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type SessionAdapter struct {
	conn *grpc.ClientConn
	context context.Context
	dispatcher protocol.ImDispatchServiceClient
}

func NewSessionAdapter(conf config.Host) *SessionAdapter {
	conn, err := grpc.Dial(conf.Address, grpc.WithInsecure())
	if err != nil {
		panic("connect to session dispatch grpc server error: " + err.Error())
	}
	ctx := context.Background()
	return &SessionAdapter{
		conn: conn,
		context: ctx,
		dispatcher: protocol.NewImDispatchServiceClient(conn),
	}
}

func (this *SessionAdapter) SetSessionInHandler(handler func(*protocol.Session)) {
	go this.on_session_in(handler)
}

func (this *SessionAdapter) SetSessionOutHandler(handler func(*protocol.Session)) {
	go this.on_session_out(handler)
}

func (this *SessionAdapter) on_session_in(cb func(*protocol.Session)) {
	cli, _ := this.dispatcher.SessionIn(context.Background(), &protocol.Empty{})
	for {
		session, err := cli.Recv()
		if err != nil {
			log.Printf("get session in error: %v", err)
			break
		}
		cb(session)
	}
}

func (this *SessionAdapter) on_session_out(cb func(*protocol.Session)) {
	cli, _ := this.dispatcher.SessionOut(context.Background(), &protocol.Empty{})
	for {
		session, err := cli.Recv()
		if err != nil {
			log.Printf("get session out error: %v", err)
			break
		}
		cb(session)
	}
}
