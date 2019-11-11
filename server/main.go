package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal("Failed to listen on tcp://127.0.0.1:5000")
	}
	defer listen.Close()
	fmt.Println("Started to listen on tcp://127.0.0.1:5000")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Could not establish the connection")
		}

		buf := make([]byte, 1024)

		go func() {
			fmt.Printf("[Remote Address]\n%s\n", conn.RemoteAddr())

			n, _ := conn.Read(buf)
			fmt.Printf("[Message]\n%s", string(buf[:n]))

			time.Sleep(3 * time.Second)

			res := fmt.Sprintf("[Response]\nHello, %s. This is server.\n", conn.RemoteAddr())
			conn.Write([]byte(res))

			conn.Close()
		}()
	}
}
