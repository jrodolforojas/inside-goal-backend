package main

import "github.com/jrodolforojas/inside-goal-backend/internal/transports"

func main() {

	server := transports.WebServer{}
	server.StartServer()
}
