package main

import (
	// router "key-management-service/biz/router"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	// router.GeneratedRegister(r)

	customizedRegister(r)
}
