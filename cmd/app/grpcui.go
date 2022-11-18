package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
)

type grpcui struct {
	router   *chi.Mux
	grpcPort string
}

func NewgGrpcui(router *chi.Mux, grpcPort string) grpcui {
	return grpcui{
		router:   router,
		grpcPort: grpcPort,
	}
}

func (g grpcui) Start(ctx context.Context) error {
	log.Println("running grpcui")
	cc, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%s", g.grpcPort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer cc.Close()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	keycloak := standalone.AddJSFile("keycloak.min.js", func() (io.ReadCloser, error) {
		path := filepath.Join(wd, "api", "static", "keycloak.min.js")
		return os.Open(path)
	})
	keycloakGrpcui := standalone.AddJSFile("keycloak-grpcui.js", func() (io.ReadCloser, error) {
		path := filepath.Join(wd, "api", "static", "keycloak-grpcui.js")
		return os.Open(path)
	})
	h, err := standalone.HandlerViaReflection(ctx, cc, fmt.Sprintf("127.0.0.1:%s", g.grpcPort), keycloak, keycloakGrpcui)
	if err != nil {
		return err
	}
	g.router.Mount("/grpcui/", http.StripPrefix("/grpcui", h))
	g.router.Get("/grpcui", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/grpcui/", http.StatusMovedPermanently)
	})
	<-ctx.Done()
	return nil
}

func (g grpcui) Compensate() {
	log.Println("grpcui stopped")
}
