package main

import (
	"dict"
	"fmt"
	"log"
	"os"

	_ "embed"
)

//go:embed dict.json
var rawDict []byte

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dict <word>")
		os.Exit(1)
	}

	dictMap, err := dict.ParseDict(rawDict, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dictMap.Query(os.Args[1]))
}
