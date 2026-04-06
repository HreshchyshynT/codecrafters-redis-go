package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Uncomment the code below to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	connection, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	for {
		// TODO: decide buffer size
		data := make([]byte, 100)
		i, err := connection.Read(data)
		if err != nil {
			log.Fatal("can't read data: ", err)
		}
		if i > 0 {
			// do not care about received content yet
			connection.Write([]byte("+PONG\r\n"))

		}
	}

}
