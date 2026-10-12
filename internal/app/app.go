package app

import (
	"net/http"

	httpadapter "github.com/betosardinha/anchovy-api/internal/adapters/http"
)

type App struct {
	Router http.Handler
}

func New() *App {
	return &App{
		Router: httpadapter.NewRouter(),
	}
}
