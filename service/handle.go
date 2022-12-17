package service

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func Handle(conn net.Conn) {
	// checkers for Lengthof chat / Name
	mutex.Lock()
	if len(users)+1 > 10 {
		conn.Write([]byte("Chat is full 10/10"))
		conn.Close()
	}
	mutex.Unlock()

	LoadLogo(conn)

	// prepartion of Users
	who := Nikcname(conn)

	newUser := User{
		conn: conn,
		name: who,
	}

	// notificaion all users about new user
	entering <- newUser
	messages <- newMessage("has joined our chat...", newUser, "")

	// First message
	t := time.Now().Format("2020-01-20 16:03:43")
	txt := fmt.Sprintf("[%s][%s]:", t, who)
	fmt.Fprintf(conn, txt)

	input := bufio.NewScanner(conn)

	for input.Scan() {
		txt = input.Text()
		if txt == "" || len(txt) < 1 {
			fmt.Fprintf(conn, "Empty text\n")
			fmt.Fprintf(conn, "[%s][%s]:", t, who)
			continue
		} else {
			if !TextChecker(txt) {
				fmt.Fprintf(conn, "Incorrect text\n")
				fmt.Fprintf(conn, "[%s][%s]:", t, who)
				continue
			} else {
				mutex.Lock()
				messages <- newMessage(txt, newUser, t)
				mutex.Unlock()

				text := fmt.Sprintf("[%s][%s]:%s\n", t, who, txt)

				mutex.Lock()
				chathistory = append(chathistory, text)
				mutex.Unlock()

				t = time.Now().Format("2020-01-20 16:03:43")
				mutex.Lock()

				mutex.Unlock()
				continue
			}
		}

	}

	// left user
	mutex.Lock()
	messages <- newMessage("has left our chat...", newUser, "")

	leaving <- newUser
	conn.Close()
	mutex.Unlock()
}

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			// mutex.Lock()

			for user, _ := range users {
				if user.name != msg.user.name {
					if msg.time != "" {
						fmt.Fprintf(user.conn, "\n[%s][%s]: %s\n", msg.time, msg.user.name, msg.msg)
					} else {
						fmt.Fprintf(user.conn, "\n%s %s\n", msg.user.name, msg.msg)
					}
					t := time.Now().Format("2020-01-20 16:03:43")
					fmt.Fprintf(user.conn, "[%s][%s]:", t, user.name)
				} else {
					if msg.time != "" {
						fmt.Fprintf(msg.user.conn, "[%s][%s]:", msg.time, msg.user.name)
					}
				}
			}
			// mutex.Unlock()
		case user := <-entering:
			mutex.Lock()
			users[user] = true
			for _, w := range chathistory {
				fmt.Fprintf(user.conn, "%s", w)
			}
			mutex.Unlock()
		case user := <-leaving:
			mutex.Lock()
			delete(users, user)
			mutex.Unlock()
		}
	}
}
