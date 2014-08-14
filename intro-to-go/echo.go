package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		fmt.Fprintf(c, "Connected!\n")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(c, c)
	}
}
