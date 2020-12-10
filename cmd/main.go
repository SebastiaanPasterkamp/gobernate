package main

import (
	"fmt"
	"gobernate"
	"gobernate/version"
	"net/http"

	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set.")
	}

	g := gobernate.New(
		port,
		version.Name,
		version.Release,
		version.Commit,
		version.BuildTime,
	)

	shutdown := g.Launch()

	g.Router.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello! Your request was processed.")
	})

	g.Ready()
	<-shutdown
}
