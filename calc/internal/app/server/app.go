package server

import (
	"context"
	"fmt"
	"github.com/d8ml/calculation_server/calc/internal/pkg/http/middleware"
	"net/http"
	"os"
)

const (
	DefaultHTTPPort = "8082"
)

type App struct {
	http.Server
	ServiceUrl string
}

func NewApp() *App {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultHTTPPort
	}
	serverUrl := fmt.Sprintf(":%s", port)

	logMux := http.NewServeMux()
	logMux.Handle("POST /api/v1/calculate", middleware.NewLogMux(Calculate))

	return &App{
		Server: http.Server{
			Addr:    serverUrl,
			Handler: logMux,
		},
	}
}

func (app *App) Start() error {
	return app.Server.ListenAndServe()
}

func (app *App) Stop(ctx context.Context) error {
	return app.Server.Shutdown(ctx)
}
