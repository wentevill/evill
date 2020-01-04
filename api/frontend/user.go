package main

import (
	pb "evill/basic/user/proto"
	"evill/internal/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type user struct {
	client pb.UserClient
}

func (u *user) name() string {
	return "evill.user:8080"
}

func (u *user) init() error {
	conn, err := grpc.Dial(u.name(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(middleware.RequestIdClient),
	)
	if err != nil {
		return err
	}
	u.client = pb.NewUserClient(conn)
	return nil
}

func (u *user) register(v ...*gin.RouterGroup) {
	v1 := v[0]
	v1.GET("/sign/up", u.signUp)
}

func (u *user) signUp(c *gin.Context) {
	params := new(pb.SignUpRequest)
	if err := c.Bind(params); err != nil {
		errResponse(c, err)
		return
	}
	if resp, err := u.client.SignUp(c, params); err != nil {
		errResponse(c, err)
		return
	} else {
		response(c, resp)
	}
}
