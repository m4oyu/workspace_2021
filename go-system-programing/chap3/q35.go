package main

import (
	"io"
	"os"
	"strings"
)

func CopyN(dest io.Writer, src io.Reader, length int64) {
	reader := io.LimitReader(src, length)
	io.Copy(dest, reader)
}

func main() {
	data := "get this header and not get this body"
	reader := strings.NewReader(data)
	io.CopyN(os.Stdout, reader, 15)
}
