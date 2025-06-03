package main

import (
	"flag"
	"fmt"
	"os/user"
)

type Command struct {
	help    *bool
	version *bool
}

func main() {
	version := "1"
	command := Command{
		help:    flag.Bool("help", false, "Display this help and exit"),
		version: flag.Bool("version", false, "Output version information and exit"),
	}
	flag.Parse()
	if !*command.help && !*command.version {
		user, err := user.Current()
		if err != nil {
			fmt.Println("Error getting the current user: ", err)
		}
		fmt.Println(user.Username)
	}
	if !*command.help && *command.version {
		fmt.Println(version)
		return
	}
	if *command.help && !*command.version {
		fmt.Println(
			"\n\nglogname [OPTION]\n\n",
			"Description:\n",
			"\tPrint the user's login name.\n\n",
			"--help display this help and exit\n\n",
			"--version output version information and exit",
		)
		return
	}
}
