package internal

import (
	"context"

	protos "github.com/mvgeny/gateway/pkg"
)

type service struct {
	protos.UnimplementedGreeterServer
}

func NewService() *service {
	return &service{}
}

func (s *service) SayHello(ctx context.Context, in *protos.HelloRequest) (*protos.HelloReply, error) {
	return &protos.HelloReply{Message: in.Name + " world"}, nil
}
