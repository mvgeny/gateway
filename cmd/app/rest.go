package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	protos "github.com/mvgeny/gateway/pkg"
	"google.golang.org/grpc"
)

type restServer struct {
	server *http.Server
}

func NewRestServer(ctx context.Context, router *chi.Mux, restPort string, grpcPort string) (restServer, error) {
	rmux := runtime.NewServeMux()
	err := protos.RegisterGreeterHandlerFromEndpoint(
		ctx,
		rmux,
		fmt.Sprintf("localhost:%s", grpcPort),
		[]grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return restServer{}, err
	}
	router.Mount("/", rmux)
	return restServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("localhost:%s", restPort),
			Handler: router,
		},
	}, nil
}

func (s restServer) Start(context.Context) error {
	log.Println("running rest server")
	return s.server.ListenAndServe()
}

func (s restServer) Compensate() {
	err := s.server.Shutdown(context.Background())
	if err == nil {
		log.Println("rest server stopped")
	}
}
