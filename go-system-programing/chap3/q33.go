package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	create, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	defer create.Close()

	zipWriter := zip.NewWriter(create)
	defer zipWriter.Close()

	data := "this is a pen"
	reader := strings.NewReader(data)

	writer, err := zipWriter.Create("newfile.txt")

	io.Copy(writer, reader)

}
