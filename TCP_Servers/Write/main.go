package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

		io.WriteString(connection, "\n Hello from Tcp Server|n")
		fmt.Fprintln(connection, "How is tour day?")
		fmt.Fprintf(connection, "%v", "well,I hope")

		connection.Close()
	}

}
