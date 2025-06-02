package main

import (
	"flag"
	"fmt"
	"runtime"
)

type Command struct {
	help    *bool
	version *bool
}

func main() {
	version := "2"
	command := Command{
		help:    flag.Bool("help", false, "Display this help and exit"),
		version: flag.Bool("version", false, "Output version information and exit"),
	}
	flag.Parse()
	if !*command.help && !*command.version {
		if runtime.GOARCH == "amd64" {
			fmt.Print("x86_64")
		}
		if runtime.GOARCH == "amd86" {
			fmt.Print("x86")
		}
	}
	if !*command.help && *command.version {
		fmt.Println(version)
		return
	}
	if *command.help && !*command.version {
		fmt.Println(
			"\n\ngarch [OPTION]\n",
			"garch OPTION\n\n",
			"Description:\n",
			"\tPrint machine architecture.\n\n",
			"--help display this help and exit\n\n",
			"--version output version information and exit",
		)
		return
	}
}
