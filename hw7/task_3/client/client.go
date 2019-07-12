package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8880")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}

}
