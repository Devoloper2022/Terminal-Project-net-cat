package service

import (
	"net"
	"sync"
)

var (
	chathistory []string
	users       = make(map[User]bool)
	entering    = make(chan User)
	leaving     = make(chan User)
	messages    = make(chan Message) // Все входящие сообщения клиента
	mutex       sync.Mutex
)

type User struct {
	name string
	conn net.Conn
}

type Message struct {
	msg  string
	user User
	time string
}
