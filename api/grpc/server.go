package grpc

import (
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/proto"
	"gitlab.com/pangold/goimb/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServer struct {
	conf config.Host
	service *service.Service
}

func NewGrpcServer(conf config.Host, impl *service.Service) *GrpcServer {
	return &GrpcServer{
		conf: conf,
		service: impl,
	}
}

func (this *GrpcServer) Run() {
	listener, err := net.Listen("tcp", this.conf.Address)
	if err != nil {
		panic("grpc listen error: " + err.Error())
	}
	log.Printf("GRPC server start serving on %s", this.conf.Address)
	srv := grpc.NewServer()
	imb.RegisterFriendServiceServer(srv, NewFriendService(this.service.FriendService))
	imb.RegisterGroupServiceServer(srv, NewGroupService(this.service.GroupService))
	reflection.Register(srv)
	if err := srv.Serve(listener); err != nil {
		panic("grpc serve " + this.conf.Address + " error: " + err.Error())
	}
}
