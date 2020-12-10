package main

import (
	"gobernate"
	"gobernate/version"

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
	g.Ready()
	<-shutdown
}
