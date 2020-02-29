package main

import (
	"log"
	"strings"
	"sync"

	"github.com/tidwall/redcon"
)

var (
	address = ":2654"
)

func main() {
	go log.Printf("started server at %s", address)

	var mu sync.RWMutex

	db := make(map[string][]byte)

	err := redcon.ListenAndServe(
		address,
		func(conn redcon.Conn, cmd redcon.Command) {

			command := strings.ToLower(string(cmd.Args[0]))

			switch command {
			default:
				conn.WriteError("not found '" + command + "'")
			case "ping":
				conn.WriteString("Pong :)")
			case "set":
				if len(cmd.Args) != 3 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}

				mu.Lock()

				db[string(cmd.Args[1])] = cmd.Args[2]

				mu.Unlock()

				conn.WriteString("OK")
			case "get":
				if len(cmd.Args) != 2 {
					conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
					return
				}
				mu.RLock()
				val, ok := db[string(cmd.Args[1])]
				mu.RUnlock()

				if !ok {
					conn.WriteNull()
				} else {
					conn.WriteBulk(val)
				}
			case "del":
				if len(cmd.Args) != 2 {
					conn.WriteString("worng number of params for this command [" + command + "]")
					return
				}

				mu.Lock()

				delete(db, string(cmd.Args[1]))

				mu.Unlock()

				conn.WriteString("OK")
			}
		},
		func(conn redcon.Conn) bool {
			// use this function to accept or deny the connection.
			log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// this is called when the connection has been closed
			log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
