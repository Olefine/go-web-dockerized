package main

import "github.com/olefine/go-dockerized-web/webserver"

func main() {
	server := webserver.NewServer("localhost", uint16(8080))

	server.Run()
}
