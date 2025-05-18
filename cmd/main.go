package main

import (
	"github.com/nosvagor/hgmx-builder/internal/server"
)

func main() {
	server.InitLogger()
	server.Start(server.New())
}
