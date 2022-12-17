package service

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

func Serve() {
	port, err := Port()
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

	go Broadcaster()
	fmt.Println(2.1)
	for {
		fmt.Println(1.1)
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("connection err: %v", err)
			continue
		}
		fmt.Println(1.2)

		go Handle(conn) // start new goroutine per connection
		fmt.Println(1.3)

	}
}

func Port() (string, error) {
	input := os.Args
	port := ""

	if len(input) == 1 {
		port = "8989"
	} else if len(input) == 2 {
		port = input[1]
	} else {
		return "", errors.New("Invalid command. To launch TCP server type \"go run . [PORT] \" or \"go run .\" ")
	}

	return port, nil
}
