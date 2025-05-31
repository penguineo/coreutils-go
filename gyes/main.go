package main

import (
	"flag"
	"fmt"
)

type Command struct {
	help    *bool
	version *bool
}

func main() {
	version := "0.0.1"
	command := Command{
		help:    flag.Bool("help", false, "Display this help and exit"),
		version: flag.Bool("version", false, "Output version information and exit"),
	}
	flag.Parse()
	if !*command.help && !*command.version && len(flag.Args()) == 0 {
		for {
			fmt.Println("y")
		}
	}
	if !*command.help && !*command.version && len(flag.Args()) != 0 {
		for {
			var args string
			for _, arg := range flag.Args() {
				args += arg + " "
			}
			fmt.Println(args)
		}
	}
	if !*command.help && *command.version {
		fmt.Println(version)
		return
	}
	if *command.help && !*command.version {
		fmt.Println(
			"\n\ngyes [STRING]\n",
			"gyes OPTION\n\n",
			"Description:\n",
			"\tRepeadly output a line with all specified STRING(s), or 'y'.\n\n",
			"--help display this help and exit\n\n",
			"--version output version information and exit",
		)
		return
	}
}
