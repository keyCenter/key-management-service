package user

import (
	"context"
	"key-management-service/biz/model/response"
)

type RegisterRequest struct {
	Username string ` form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginRequest struct {
	Username string ` form:"username" json:"username"`
	Password string ` form:"password" json:"password"`
}

type InfoRequest struct {
	Username string `json:"Username" path:"username"`
}

type UserService interface {
	Register(ctx context.Context, req *RegisterRequest) (r *response.Result, err error)

	Login(ctx context.Context, req *LoginRequest) (r **response.Result, err error)

	Info(ctx context.Context, req *InfoRequest) (r **response.Result, err error)
}
