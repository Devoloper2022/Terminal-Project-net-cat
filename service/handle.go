package service

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func Handle(conn net.Conn) {
	if len(users) > 9 {
		conn.Write([]byte("Chat is full 10/10"))
		conn.Close()
		return
	}

	LoadLogo(conn)
	name := Nikcname(conn)

	mutex.Lock()
	users[name] = conn
	mutex.Unlock()

	mutex.Lock()
	for _, v := range chathistory {
		fmt.Fprintln(conn, string(v))
	}
	mutex.Unlock()

	t := time.Now().Format("2020-01-20 16:03:43")

	mutex.Lock()
	join <- newMessage("has joined our chat...", name, t)
	mutex.Unlock()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		t := time.Now().Format("2020-01-20 16:03:43")

		mutex.Lock()
		messages <- newMessage(input.Text(), name, t)
		mutex.Unlock()
	}
	mutex.Lock()
	delete(users, name)
	leaving <- newMessage("has left our chat...", name, time.Now().Format(""))
	mutex.Unlock()
	conn.Close()
}

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for user, c := range users {
				fmt.Fprintf(c, "\n %s %s\n", msg.userName, msg.msg)
			}
			mutex.Unlock()
		case msg := <-join:
			mutex.Lock()
			for user, c := range users {
				if msg.userName != user {
					fmt.Fprintf(c, "\n %s %s\n", msg.userName, msg.msg)
				}
			}
			mutex.Unlock()
		case msg := <-leaving:
			mutex.Lock()
			for user, c := range users {
				if msg.userName != user {
					fmt.Fprintf(c, "\n %s %s\n", msg.userName, msg.msg)
				}
			}
			mutex.Unlock()
		}
	}
}
