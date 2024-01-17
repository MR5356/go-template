package main

import (
	"github.com/MR5356/go-template/config"
	"github.com/MR5356/go-template/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New(config.WithPort(1002))

	srv, err := server.New(cfg)
	if err != nil {
		logrus.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Failed to run server: %v", err)
	}
}
