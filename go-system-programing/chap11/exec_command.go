package main

import (
	"fmt"
	"os"
	"os/exec"
)

// go run exex_commend.go sleep 1
func main() {
	if len(os.Args) == 1 {
		return
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("PID: %d\n", state.Pid())
	fmt.Printf("System: %v\n", state.SystemTime())
	fmt.Printf("User: %v\n", state.UserTime())
}
