package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 9.5.1 ディレクトリのパスとファイル名とを連結
	fmt.Printf("temp file path: %s\n", filepath.Join(os.TempDir(), "temp.txt"))

	// 9.5.2 パスの分割
	dir, name := filepath.Split(os.Getenv("GOPATH"))
	fmt.Printf("dir: %s, name: %s\n", dir, name)

	// 9.5.4 パスのクリーン化
	// ..の削除や絶対・相対パスを相互変換
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))

	// 9.5.6 ファイル名検索
	files, err := filepath.Glob("./*.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	// 9.5.3 複数パスの分割
	if len(os.Args) == 1 {
		fmt.Printf("%s [exec file name]\n", os.Args[0])
		os.Exit(1)
	}
	for _, path := range filepath.SplitList(os.Getenv("PATH")) {
		execpath := filepath.Join(path, os.Args[1])
		_, err := os.Stat(execpath)
		if !os.IsNotExist(err) {
			fmt.Println(execpath)
			return
		}
	}
	os.Exit(1)

}
