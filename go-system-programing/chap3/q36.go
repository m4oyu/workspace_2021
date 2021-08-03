package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

var (
	computer   = strings.NewReader("COMPUTER")
	system     = strings.NewReader("SYSTEM")
	programing = strings.NewReader("PROGRAMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programing, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i := io.NewSectionReader(programing, 7, 1)
	i2 := io.NewSectionReader(programing, 7, 1)

	// ex 1 : recommend
	//stream = io.MultiReader(a, s, c, i, i2)

	// ex 2 : not good
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	io.Copy(writer, a)
	io.Copy(writer, s)
	io.Copy(writer, c)
	io.Copy(writer, i)
	io.Copy(writer, i2)
	stream = bufio.NewReader(&buffer)

	io.Copy(os.Stdout, stream)
}
