package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	hello "micro_example_project/app/hello/service"
	"micro_example_project/config"
	HelloPb "micro_example_project/idl/hello"
)

func main() {
	config.Init()

	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	//定义服务
	service := micro.NewService(
		micro.Name("HelloService"),
		micro.Address(config.HelloServiceAddress),
		micro.Registry(etcdReg),
	)

	//服务初始化
	service.Init()

	// 注册handler
	err := HelloPb.RegisterGreeterHandler(service.Server(), hello.GetHelloSrv())
	if err != nil {
		return
	}

	//启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
