package main

import (
	"log"
	"net/http"

	"github.com/betosardinha/anchovy-api/internal/app"
	"github.com/betosardinha/anchovy-api/internal/config"
)

func main() {
	a := app.New()
	c := config.Load()

	s := &http.Server{
		Addr:    ":" + c.Port,
		Handler: a.Router,
	}

	log.Println("listening on " + c.Port)

	s.ListenAndServe()
}
