package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("listen tick server at 224.0.0.1:9999")
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddr, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("server %v\n", remoteAddr)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
