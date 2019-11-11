package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal("Failed to listen on tcp://127.0.0.1:5000")
	}
	defer conn.Close()

	msg := fmt.Sprintf("Hello, %s. This is client.\n", conn.RemoteAddr())
	conn.Write([]byte(msg))

	res := make([]byte, 1024)
	n, _ := conn.Read(res)
	fmt.Println(string(res[:n]))
}
