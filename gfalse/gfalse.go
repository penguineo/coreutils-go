package main

import (
	"fmt"
	"os"
)

func main() {
	version := "1"
	for _, args := range os.Args[1:] {
		switch args {
		case "--help":
			fmt.Println(
				"\n\ngfalse [ignored command line arguements]\n",
				"gfalse OPTION\n\n",
				"Description:\n",
				"\tExit with a status code indicating success.\n\n",
				"--help display this help and exit\n\n",
				"--version output version information and exit\n\n",
				"Your shell may have its own version of true, which usually supersedes the version described here.",
				"Please refer to your shell's documentation for details about options it supports.",
			)
			os.Exit(1)
		case "--version":
			fmt.Println("gfalse Version:", version)
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
