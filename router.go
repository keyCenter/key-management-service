package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"key-management-service/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
