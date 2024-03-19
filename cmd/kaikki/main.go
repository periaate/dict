package main

import (
	"fmt"
	"os"

	"github.com/periaate/dict"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: kaikki <input> <output>")
		os.Exit(1)
	}

	b, err := dict.FormatKaikki(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
