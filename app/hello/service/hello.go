package service

import (
	"context"
	HelloPb "micro_example_project/idl/hello"
	"sync"
)

type HelloSrv struct {
}

var HelloSrvIns *HelloSrv
var HelloSrvOnce sync.Once

func GetHelloSrv() *HelloSrv {
	HelloSrvOnce.Do(func() {
		HelloSrvIns = &HelloSrv{}
	})
	return HelloSrvIns
}

func (h *HelloSrv) Hello(ctx context.Context, in *HelloPb.Request, out *HelloPb.Response) error {
	out.Greeting = "Hello " + in.Name
	return nil
}
