package main

import (
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Hanlder(r)
}
