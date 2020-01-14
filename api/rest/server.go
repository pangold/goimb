package rest

import (
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/service"
)

type RestServer struct {
	router *Router
}

func NewRestServer(conf config.Host, impl *service.Service) *RestServer {
	return &RestServer{
		router: NewRouter(conf, impl),
	}
}

func (this *RestServer) Run() {
	this.router.Run()
}
