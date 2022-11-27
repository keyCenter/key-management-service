package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(
		server.WithHostPorts("127.0.0.1:8000"),
	)

	register(h)
	h.Spin()
}
