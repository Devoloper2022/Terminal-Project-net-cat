package service

import (
	"net"
	"sync"
)

type client chan<- Message // Канал исходящих сообщений

var (
	chathistory []string
	users       = make(map[string]net.Conn)
	join        = make(chan Message)
	leaving     = make(chan Message)
	messages    = make(chan Message) // Все входящие сообщения клиента
	mutex       sync.Mutex
)

type Message struct {
	msg      string
	userName string
	time     string
}
