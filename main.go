package main

import (
	"goown-vpn/server"
	"os"	
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := server.NewServer()
	server.Run(":" + port)
}