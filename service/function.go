package service

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func LoadLogo(conn net.Conn) {
	file, err := os.ReadFile("logo.txt")
	if err != nil {
		fmt.Printf("couldn't read this file")
		return
	}

	logo := string(file)
	conn.Write([]byte(logo + "\n"))
}

func Nikcname(conn net.Conn) string {
	// get name
	conn.Write([]byte("[Enter your name]: "))
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Fprintln(conn, err)
		return ""
	}

	// delete space leading and trailing
	name := strings.TrimSpace(data)

	// check for nil
	if name == "" || len(name) == 0 {
		fmt.Fprintln(conn, "Incorrect input")
		return Nikcname(conn)
	}

	for _, w := range name {
		if w < 32 || w > 127 {
			fmt.Fprintln(conn, "Incorrect input")
			return Nikcname(conn)
		}
	}
	for user, _ := range users {
		if user == name {
			fmt.Fprintf(conn, "User already exist\n")
			return Nikcname(conn)
		}
	}
	return name
}

func newMessage(msg string, user string, time string) Message {
	return Message{
		msg:      msg,
		userName: user,
		time:     time,
	}
}
