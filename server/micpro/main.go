package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/protoc-gen-micro/plugin/micro"
	"micpro/handler"
	micpro "micpro/proto/micpro"
)

func main() {
	consulRegis := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.micpro"),
		micro.Version("latest"),
		micro.Registry(consulRegis),
		micro.Address("127.0.0.1:10345"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	micpro.RegisterMicproHandler(service.Server(), new(handler.Micpro))
	fmt.Println("微服务已运行......")
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
