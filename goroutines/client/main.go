package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/LearnGoFarsi/go-concurrency/goroutines"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:"+strconv.FormatInt(goroutines.PORT, 10))
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, conn); err != nil {
		fmt.Println(err)
	}
}
