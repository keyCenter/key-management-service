package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"key-management-service/biz/router"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
