package main

import (
	"fmt"
	"os"
	"testing"
)

func f() {
	fmt.Println("do random actions")
}
func Test_f(t *testing.T) {
	f()
}

func TestMain(m *testing.M) {
	fmt.Println("before action")
	status := m.Run()
	fmt.Println("after action")

	os.Exit(status)
}
