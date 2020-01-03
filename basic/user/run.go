package main

import (
	user "evill/basic/user/proto"
	"evill/einit"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	assembly, err := einit.Init(einit.Mysql|einit.Log, "./config/user.yml")
	if err != nil {
		panic(err)
	}

	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 继续处理请求
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			requestIds := md["request_id"]
			if len(requestIds) != 0 {
				ctx = context.WithValue(ctx, "request_id", requestIds[0])
			}
			return handler(ctx, req)
		}
		return handler(ctx, req)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)

	reflection.Register(s)
	user.RegisterUserServer(s, &User{db: assembly.Mysql()})

	lis, err := net.Listen("tcp", einit.GetConfig().Port)
	if err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
	log.Info("server start success")
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
}
