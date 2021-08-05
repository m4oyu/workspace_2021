package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	request, err := http.NewRequest(
		"GET",
		"http://localhost:8080",
		nil)
	if err != nil {
		panic(err)
	}

	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	response, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	dumpResponse, err := httputil.DumpResponse(response, false)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dumpResponse))

	if len(response.TransferEncoding) < 1 || response.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		parseInt, err := strconv.ParseInt(string(bytes[:len(bytes)-2]), 16, 64)
		if parseInt == 0 {
			break
		}
		if err != nil {
			panic(err)
		}

		line := make([]byte, int(parseInt))
		io.ReadFull(reader, line)
		reader.Discard(2)
		fmt.Printf("	%d bytes: %s\n", parseInt, string(line))
	}

}
