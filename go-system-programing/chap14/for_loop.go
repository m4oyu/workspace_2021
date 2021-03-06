package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, task := range tasks {
		go func(task string) {
			// goroutineが起動するときにはfor分が回り切って
			// すべて"cpack"が出力される
			fmt.Println(task)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
