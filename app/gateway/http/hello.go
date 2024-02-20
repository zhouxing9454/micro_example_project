package http

import (
	"github.com/gin-gonic/gin"
	"micro_example_project/app/gateway/rpc"
	"micro_example_project/idl/hello"
	"net/http"
)

func HelloHandler(ctx *gin.Context) {
	var req hello.Request
	req.Name = ctx.Query("name")
	var res *hello.Response
	res, _ = rpc.Hello(ctx, &req)
	ctx.JSON(http.StatusOK, gin.H{
		"status_msg": res.Greeting,
	})
}
