package service

import (
	"errors"
	"os"
)

func Serve() {
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
