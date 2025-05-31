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
				"\n\ngtrue [ignored command line arguements]\n",
				"gtrue OPTION\n\n",
				"Description:\n",
				"\tExit with a status code indicating success.\n\n",
				"--help display this help and exit\n\n",
				"--version output version information and exit\n\n",
				"Your shell may have its own version of true, which usually supersedes the version described here.",
				"Please refer to your shell's documentation for details about options it supports.",
			)
			os.Exit(0)
		case "--version":
			fmt.Println("gtrue Version:", version)
			os.Exit(0)
		default:
			os.Exit(0)
		}
	}
}
