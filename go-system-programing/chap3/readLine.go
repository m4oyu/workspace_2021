package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `line1
line2 apple
line3`

//func main() {
//	reader := bufio.NewReader(strings.NewReader(source))
//	for {
//		readString, err := reader.ReadString('\n')
//		fmt.Printf("%#v\n", readString)
//		if err == io.EOF {
//			break
//		}
//	}
//}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	// 文字区切り
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}

// scanner.Split()で区切り文字指定
