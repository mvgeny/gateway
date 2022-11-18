package app

import (
	"context"
	"fmt"
	"log"
	"net"

	protos "github.com/mvgeny/gateway/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	server *grpc.Server
	lis    net.Listener
}

func NewGrpcServer(srv protos.GreeterServer, grpcPort string) (grpcServer, error) {
	s := grpc.NewServer()
	reflection.Register(s)
	protos.RegisterGreeterServer(s, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", grpcPort))
	if err != nil {
		return grpcServer{}, err
	}
	return grpcServer{
		server: s,
		lis:    lis,
	}, nil
}

func (s grpcServer) Start(context.Context) error {
	log.Println("running grpc server")
	return s.server.Serve(s.lis)
}

func (s grpcServer) Compensate() {
	s.server.GracefulStop()
	log.Println("grpc server stopped")
}
