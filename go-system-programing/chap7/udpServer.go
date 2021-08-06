package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server is running at localhost:8080")
	conn, err := net.ListenPacket("udp4", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		conn.WriteTo([]byte("Hello from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
