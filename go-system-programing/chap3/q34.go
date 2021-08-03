package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	// q33.go
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	data := "this is a pen"
	reader := strings.NewReader(data)
	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(writer, reader)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
