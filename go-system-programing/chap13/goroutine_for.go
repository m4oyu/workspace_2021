package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutineが起動するときにはfor分が回り切って
			// すべて"cpack"が出力される
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
