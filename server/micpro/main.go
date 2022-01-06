package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"micpro/handler"
	"micpro/subscriber"

	micpro "micpro/proto/micpro"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.micpro"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	micpro.RegisterMicproHandler(service.Server(), new(handler.Micpro))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.micpro", service.Server(), new(subscriber.Micpro))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.micpro", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
