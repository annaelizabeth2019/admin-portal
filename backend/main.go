package main

import (
	"admin-portal/pkg/config"
	"admin-portal/pkg/logger"
)

func main() {
	if err := config.Setup(); err != nil {
		logger.Fatalf("config.Setup() error: %s", err)
	}
}
