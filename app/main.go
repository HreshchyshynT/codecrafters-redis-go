package main

import (
	"errors"
	"fmt"
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

	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go listenConnection(connection)
	}
}

func listenConnection(c net.Conn) {
	defer c.Close()
	data := make([]byte, 4<<10) // 4 KB buffer
	for {
		// read [0, i] bytes, (i, len(data)) - are stale
		i, err := c.Read(data)

		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				fmt.Println("Connection closed")
			} else {
				fmt.Println("can't read data: ", err.Error())
			}
			break
		}

		if i > 0 {
			// do not care about received content yet
			c.Write([]byte("+PONG\r\n"))
		}
	}
}
