package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"key-management-service/biz/model/response"
	"key-management-service/biz/model/user"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, response.Failed(400, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, response.Success(fmt.Sprintf("%s", req)))
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)

	if err != nil {
		c.JSON(consts.StatusOK, response.Failed(400, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, response.Success(fmt.Sprintf("%s", req)))
}

// Info .
// @router /user/:username [GET]
func Info(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.InfoRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, response.Failed(400, err.Error()))
		return
	}

	c.JSON(consts.StatusOK, response.Success(fmt.Sprintf("%s", req)))
}
