package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8880")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			log.Print(err)
			continue
		}
		go message(conn)
		go handleConn(conn)

	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func message(c net.Conn) {
	defer c.Close()
	for {
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print(message)

	}
}
