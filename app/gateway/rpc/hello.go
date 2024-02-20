package rpc

import (
	"context"
	"micro_example_project/idl/hello"
)

func Hello(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	res, err = HelloService.Hello(ctx, req)
	if err != nil {
		return
	}
	return
}
