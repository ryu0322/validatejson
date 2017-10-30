package main

import (
	"fmt"
	"os"

	"github.com/ryu0322/validatejson/jsons"

	"github.com/ryu0322/validatejson/reader"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("引数が不足しています。読み込みファイルを指定してください。\n")
		return
	}

	rows := reader.Reader(os.Args[1])

	maps := jsons.CreateActionGroup(rows)

	jsons.NewJson(maps)
}
