package main

import (
	"github.com/nosvagor/hgmx-builder/internal/server"
	"github.com/nosvagor/hgmx-builder/internal/shared"
)

func main() {
	cfg := shared.Load()
	server.InitLogger(cfg)
	server.Start(server.New())
}
