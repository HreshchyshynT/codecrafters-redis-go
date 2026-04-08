package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

func main() {
	// Uncomment the code below to pass the first stage
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
	decoder := resp.NewDecoder(c)
	for {
		value, err := decoder.Decode()
		if err != nil {
			if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
				log.Println("Connection closed")
			} else {
				c.Write(fmt.Appendf(nil, "can't read data: %v\n", err.Error()))
			}
			break
		}

		switch value.Typ {
		case resp.Array:
			c.Write([]byte("+array received\r\n"))
		default:
			c.Write([]byte("+PONG\r\n"))
		}
	}
}
