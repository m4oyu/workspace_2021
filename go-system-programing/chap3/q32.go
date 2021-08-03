package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	reader := rand.Reader
	limitReader := io.LimitReader(reader, 1024)

	create, err := os.Create("q32Bin")
	if err != nil {
		panic(err)
	}
	defer create.Close()

	io.Copy(create, limitReader)
}
