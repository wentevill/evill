package main

import (
	pb "evill/basic/user/proto"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type user struct {
	client pb.UserClient
}

func (u *user) name() string {
	return "evill.user:8080"
}

func (u *user) init() error {
	var interceptor grpc.UnaryClientInterceptor
	interceptor = func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md := metadata.Pairs("request_id", ctx.Value("request_id").(string))
		ctx = metadata.NewOutgoingContext(ctx, md)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	conn, err := grpc.Dial(u.name(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor),
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
