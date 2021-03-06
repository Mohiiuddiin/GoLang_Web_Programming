package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	err := connection.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "I heard you say: %s\n", line)
	}

	defer connection.Close()

	// now we get here
	// the connection will time out
	// that breaks us out of the scanner loop
	fmt.Println("***Code Go Here***")
}
