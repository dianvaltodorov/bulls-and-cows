package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const tries = 10
const intro = `Cows and Bulls
Guess the number your opponent chose with minimal tries. It is guaranteed that:
1. All digits of the number are different. 
2. The opponent number does not start with a 0.
3. A correctly guessed digit but not in the correct place is a cow. 
4. A correctly guessed digit in the correct place is a bull.
`

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte(intro))
	if err != nil {
		log.Fatal(err)
		return
	}

	num := genNumber()

	for index := 0; index < tries; index++ {
		guess, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			log.Print(err)
			return
		}

		guess = bytes.TrimSpace(guess)

		bulls, cows, err := bullsAndCows(num, guess)

		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			if bulls == allowedLength {
				_, err = conn.Write([]byte("You guessed the number"))
				if err != nil {
					log.Print(err)
					return
				}
				return
			}
			_, err = conn.Write([]byte(fmt.Sprintf("Bulls: %d, Cows: %d,\n", bulls, cows) + "\n"))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}
}

func genNumber() []byte {
	pat := make([]byte, allowedLength)
	rand.Seed(time.Now().Unix())
	r := rand.Perm(9)
	offset := 0

	for r[0] == 0 { // yes, kind of hacky, no guarantees for time complexity here
		r = rand.Perm(9)
	}

	for i := range pat {
		pat[i] = '0' + byte(r[i+offset])
	}

	return pat
}
