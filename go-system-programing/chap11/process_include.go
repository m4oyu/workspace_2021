package main

import (
	"fmt"
	"os"
)

func main() {

	// 1 実行ファイル名
	// go run ではなくシェルから.exeを起動する
	executable, _ := os.Executable()
	fmt.Printf("実行ファイル    : %s\n", os.Args[0])
	fmt.Printf("実行ファイルpath: %s\n", executable)

	// 2 process id
	fmt.Printf("process id:        %d\n", os.Getpid())
	fmt.Printf("parent process id: %d\n", os.Getppid())

	// 3. process group, session group
	// windowsでは実装されていない
	// syscall.Getsid()

	// 4. userID, groupID
	fmt.Printf("user id:      %d\n", os.Getuid())
	fmt.Printf("group id:     %d\n", os.Getgid())
	getgroups, _ := os.Getgroups()
	fmt.Printf("sub group id: %d\n", getgroups)

	// 5. work directory
	wd, _ := os.Getwd()
	fmt.Printf("work dir: %s\n", wd)
}
