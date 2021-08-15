package main

import (
	"fmt"
	"time"
)

func main() {
	tickTimer()
}

func tickTimer() {
	// 5秒ごとに出力
	fmt.Println("waiting 5 seconds")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now: ", now)
	}
}

func channelTimer() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("done")
}

func sleepTimer() {
	fmt.Println("waiting 5 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("done")
}
