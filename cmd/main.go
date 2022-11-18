package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	_ "github.com/mvgeny/gateway/api/swaggerdist"
	"github.com/mvgeny/gateway/cmd/app"
	"github.com/mvgeny/gateway/internal"
	"github.com/rakyll/statik/fs"
)

func main() {
	ctx := context.Background()
	router := chi.NewRouter()
	err := addSwagger(router)
	fatalIfErr(err)
	grpcServer, err := app.NewGrpcServer(internal.NewService(), "8092")
	fatalIfErr(err)
	restServer, err := app.NewRestServer(ctx, router, "8082", "8092")
	fatalIfErr(err)
	grpcui := app.NewgGrpcui(router, "8092")
	err = app.Run(ctx, grpcServer, restServer, grpcui)
	fatalIfErr(err)
}

func addSwagger(router *chi.Mux) error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	staticServer := http.FileServer(statikFS)
	router.Mount("/swaggerui", http.StripPrefix("/swaggerui", staticServer))
	router.Get("/swaggerui", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swaggerui/", http.StatusMovedPermanently)
	})
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "api", "static")
	router.Mount("/static", http.StripPrefix("/static", http.FileServer(http.Dir(path))))
	return nil
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
