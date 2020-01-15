package main

import (
	"os"
	"goown-vpn/cmd/server"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := server.NewServer()
	server.Run(":" + port)
}