package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buff := make([]byte, 4096)
	_, err = bufio.NewReader(conn).Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Print(string(buff))

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Guess: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(conn, text+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(message)
	}
}
