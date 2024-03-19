package main

import (
	"fmt"
	"os"

	"github.com/periaate/dict"

	_ "embed"
)

//go:embed dict.json
var rawDict []byte

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dict <word>")
		os.Exit(1)
	}
	print(dict.QueryRaw(rawDict, os.Args[1], nil))
}
