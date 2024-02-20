package router

import (
	"github.com/gin-gonic/gin"
	"micro_example_project/app/gateway/http"
	"micro_example_project/config"
)

func NewRouter() *gin.Engine {
	config.Init()
	r := gin.Default()
	r.GET("/hello", http.HelloHandler)
	return r
}
