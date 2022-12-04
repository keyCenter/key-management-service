package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "key-management-service/biz/router/user"
)

func GeneratedRegister(r *server.Hertz) {
	user.Register(r)
}
