package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const name = "request_id"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(name, generateId())
	}
}

func RequestIdServer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 继续处理请求
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		requestIds := md[name]
		if len(requestIds) != 0 {
			ctx = context.WithValue(ctx, name, requestIds[0])
		}
	}
	return handler(ctx, req)
}

func RequestIdClient(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	md := metadata.Pairs(name, ctx.Value(name).(string))
	ctx = metadata.NewOutgoingContext(ctx, md)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func generateId() string {
	u := uuid.NewV1()
	return u.String()
}
