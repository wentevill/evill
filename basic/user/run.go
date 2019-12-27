package main

import (
	user "evill/basic/user/proto"
	"evill/einit"
	"google.golang.org/grpc"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
)

func main() {
	assembly, err := einit.Init(einit.Mysql|einit.Log, "./einit/config.yml")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_opentracing.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor()),
		))

	user.RegisterUserServer(s, &User{db: assembly.Mysql()})

	reflection.Register(s)

	lis, err := net.Listen("tcp", einit.GetConfig().Port)
	if err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
	log.Info("server start success")
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to listen: " + err.Error())
	}
}
