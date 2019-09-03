package main

import (
	"context"
	"fmt"

	proto "demo/MyMircoDemo/services/pb"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("hello-world.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewHelloWorldService("hello-world", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
