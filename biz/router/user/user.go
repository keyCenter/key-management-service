package User

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"key-management-service/biz/handler/user"
)

func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/:username", append(_infoMw(), user.Info)...)
		_user.POST("/login", append(_loginMw(), user.Login)...)
		_user.POST("/register", append(_registerMw(), user.Register)...)
	}
}
