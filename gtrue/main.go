package main

import (
	"fmt"
	"os"
)

func main() {
	for _, args := range os.Args[1:] {
		switch args {
		default:
			os.Exit(0)
		}
	}
}
