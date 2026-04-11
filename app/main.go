package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
	"github.com/codecrafters-io/redis-starter-go/internal/store"
)

var cache *store.Store

func main() {
	// Uncomment the code below to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	cache = store.NewStore()

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

					// TODO: should we validate key type here?
				case "set":

					// TODO: refactor this
					var expireIn time.Duration
					if len(value.Array) > 3 && strings.ToLower(value.Array[3].String) == "px" {
						var i int
						if value.Array[4].Typ == resp.Integer {
							i = value.Array[4].Integer
						} else {
							i, _ = strconv.Atoi(value.Array[4].String)
						}
						expireIn = time.Duration(i) * time.Millisecond
					}

					cache.Put(store.Key(value.Array[1].String), store.NewData(value.Array[2], expireIn))
					encoder.Encode(resp.NewString("OK"))
				case "get":
					d, ok := cache.Get(store.Key(value.Array[1].String))
					if !ok {
						encoder.Encode(resp.NewNullBulkString())
						continue
					}
					err = encoder.Encode(d.Value)
				}
			}
			if err != nil {
				log.Println("error encoding response: ", err.Error())
			}
		default:
		}
	}
}
