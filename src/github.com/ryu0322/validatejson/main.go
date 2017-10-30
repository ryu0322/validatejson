package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("引数が不足しています。読み込みファイルを指定してください。\n")
		return
	}
}
