package rpc

import (
	"go-micro.dev/v4"
	"micro_example_project/idl/hello"
)

var (
	HelloService hello.GreeterService
)

func InitRPC() {
	HelloMicroService := micro.NewService(micro.Name("HelloService.client"))
	helloService := hello.NewGreeterService("HelloService", HelloMicroService.Client())
	HelloService = helloService
}
