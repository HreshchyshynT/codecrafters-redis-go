package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

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
	encoder := resp.NewEncoder(c)
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
			if len(value.Array) == 0 {
				continue
			}
			commandValue := value.Array[0]
			if commandValue.Typ == resp.BulkString {
				switch strings.ToLower(commandValue.String) {
				case "ping":
					err = encoder.Encode(resp.NewString("PONG"))
				case "echo":
					err = encoder.Encode(value.Array[1])
				}
			}
			if err != nil {
				log.Println("error encoding response: ", err.Error())
			}
		default:
		}
	}
}
