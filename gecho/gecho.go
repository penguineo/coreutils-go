package main

import (
	"flag"
	"fmt"
	"strings"
)

type Command struct {
	n       *bool
	e       *bool
	E       *bool
	help    *bool
	version *bool
}

func main() {
	version := "1"
	command := Command{
		n:       flag.Bool("n", false, "do not output the trailing newline"),
		e:       flag.Bool("e", false, "enable interpretation of backslash escapes"),
		E:       flag.Bool("E", true, "disable interpretation of backslash escapes"),
		help:    flag.Bool("help", false, "display this help and exit"),
		version: flag.Bool("version", false, "output version information and exit"),
	}
	flag.Parse()
	if *command.e {
		*command.E = false
	}
	args := strings.Join(flag.Args(), " ")
	if *command.e {
		args = interpretEscapes(args)
	}
	if *command.n {
		fmt.Print(args)
		return
	} else {
		fmt.Println(args)
	}
	if *command.version && !*command.n && !*command.e && !*command.help {
		fmt.Println("gecho Version:", version)
		return
	}
	if *command.help && !*command.n && !*command.e && !*command.version {
		fmt.Println(
			"\n\ngecho [SHORT-OPTION]... [STRING]...\n",
			"gecho LONG-OPTION\n\n",
			"Description:\n",
			"\tEcho the STRING(s) to standard output.\n\n",
			"\t-n \tdo not output the trailing newline\nn",
			"\t-e \tenable interpretation of backslash escapes\n\n",
			"\t-E \t disable  interpretation of backslash escapes (default)\n\n",
			"\t--help \t display this help and exit\n\n",
			"\t--version \t output version information and exit\n\n",
			"Your shell may have its own version of true, which usually supersedes the version described here.",
			"Please refer to your shell's documentation for details about options it supports.",
		)
		return
	}
}

func interpretEscapes(s string) string {
	return strings.NewReplacer(
		`\\`, `\`,
		`\n`, "\n",
		`\t`, "\t",
		`\r`, "\r",
		`\b`, "\b",
		`\f`, "\f",
		`\a`, "\a",
		`\"`, "\"",
		`\'`, "'",
	).Replace(s)
}
