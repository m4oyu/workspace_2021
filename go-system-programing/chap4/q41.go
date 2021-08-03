package main

import (
	"fmt"
	"time"
)

func main() {
	after := time.After(10 * time.Second)
	fmt.Println("timer start")
	time := <-after
	fmt.Printf("current time is : %v\n", time)
}
