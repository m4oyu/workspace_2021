package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"あああああああああああああああああああああああああああああああああ",
	"iiiiiiiiiiiiiiいいいいいいいいいいいいいいいいいいいいいいいいい",
	"うううううううううううううううううううううううううううううううううううu",
	"eeeeeeeeeeeeeeeeえええええええええええええええ",
	"ooooooooooooooおおおおおおおおおおおおおおおおおおおおおおおおおおおおお",
	"kakakakakakakkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk",
}

func chunkProcessSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()

	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dumpRequest, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dumpRequest))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("server is running at localhost:8080")
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go chunkProcessSession(conn)
	}
}
