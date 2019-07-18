package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	// Добавил переменную с текущим временем
	t := time.Now().Format("15:04:05")
	who := conn.RemoteAddr().String()

	ch <- "You are " + who
	messages <- t + ": " + who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- t + ": " + who + ": " + input.Text()
	}
	leaving <- ch
	messages <- t + ": " + who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

/*
11ap.abobkov@11AP-PC-108 MINGW64 ~/go/src/GoCourse/hw8/chat/client (homework8)
$ go run chat-client.go
You are 127.0.0.1:27935
hi
15:48:30: 127.0.0.1:27935: hi
15:49:01: 127.0.0.1:27945 has arrived
15:49:01: 127.0.0.1:27945: hi all
*/
