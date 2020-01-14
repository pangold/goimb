package api

import (
	"gitlab.com/pangold/goimb/api/grpc"
	"gitlab.com/pangold/goimb/api/rest"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/service"
)

type Api interface {
	Run()
}

type Server struct {
	apis []Api
}

func NewApiServer(conf config.ApiConfig, impl *service.Service) *Server {
	server := &Server{}
	for _, protocol := range conf.Protocols {
		if protocol == "grpc" {
			server.apis = append(server.apis, grpc.NewGrpcServer(conf.Grpc, impl))
		} else if protocol == "http" {
			server.apis = append(server.apis, rest.NewRestServer(conf.Http, impl))
		}
	}
	return server
}

func (this *Server) Run() {
	if len(this.apis) == 0 {
		panic("invalid api server")
	}
	for i := 1; i < len(this.apis); i++ {
		go this.apis[i].Run()
	}
	this.apis[0].Run()
}
