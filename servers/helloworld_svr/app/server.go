package main

import (
	"context"
	"fmt"

	proto "demo/MyMircoDemo/services/pb"
	"github.com/micro/go-micro"
)

type HelloWorld struct{}

func (g *HelloWorld) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("hello-world"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterHelloWorldHandler(service.Server(), new(HelloWorld))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
