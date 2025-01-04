package main

import "new-go-project/cmd/server"

func main() {
	serve := server.NewServer()
	serve.Start()
}
