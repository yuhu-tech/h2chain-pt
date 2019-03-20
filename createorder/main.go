package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"H2chainProject/createorder/handler"
	"H2chainProject/createorder/client"

	example "H2chainProject/createorder/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.order_api.createorder"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
