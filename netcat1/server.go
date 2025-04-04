package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("Server is running on port 8000...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte("Hello from server\n"))
		if err != nil {
			log.Println("Send error:", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
