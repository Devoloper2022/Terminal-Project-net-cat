package main

import (
	"fmt"
	"log"
	"net"
	"sonik/service"
)

func main() {
	port, err := service.Port()
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := net.Listen("tcp", "localhost:"+port) // localhost:8080
	if err != nil {
		log.Fatalf("could not start chat: %v", err)
	}

	defer server.Close()

	fmt.Println("Listening on the port :" + port)

	go service.Broadcaster()
	fmt.Println(2.1)
	for {
		fmt.Println(1.1)
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("connection err: %v", err)
			continue
		}
		fmt.Println(1.2)

		go service.Handle(conn) // start new goroutine per connection
		fmt.Println(1.3)

	}
}
