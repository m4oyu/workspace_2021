package main

import (
	"flag"
	"io"
	"os"
)

var input string
var output string

func init() {
	flag.StringVar(&input, "in", "input.txt", "-in input-file")
	flag.StringVar(&output, "out", "output.txt", "-out output-file")
}

func main() {
	flag.Parse()
	open, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer open.Close()
	create, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer create.Close()

	io.Copy(create, open)

}
