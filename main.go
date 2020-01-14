package main

import (
	"gitlab.com/pangold/goimb/api"
	"gitlab.com/pangold/goimb/config"
	"gitlab.com/pangold/goimb/front"
	"gitlab.com/pangold/goimb/service"
)

func main() {
	conf := config.NewYaml("config/config.yml").ReadConfig()
	// FIXME: Front Adapter Independence
	frontApi := front.NewFrontApi(conf.Front)
	services := service.NewService(conf.Service, frontApi)
	server := api.NewApiServer(conf.Api, services)
	server.Run()
}
