package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/LearnGoFarsi/go-concurrency/goroutines"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:"+strconv.FormatInt(goroutines.PORT, 10))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Conneciton accepted")
		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	for {
		if _, err := io.WriteString(conn, time.Now().Format(time.Stamp+"\n")); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Second)
	}
}
