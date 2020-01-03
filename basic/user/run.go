package main

import (
	user "evill/basic/user/proto"
	"evill/einit"
	"evill/internal/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	assembly, err := einit.Init(einit.Mysql|einit.Log, "./config/user.yml")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(middleware.RequestIdServer))
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
