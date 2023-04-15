package main

import (
	"log"

	"github.com/dlph/go-play/http/server"

	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	return server.ListenAndServe(server.WithLogger(l))
}
