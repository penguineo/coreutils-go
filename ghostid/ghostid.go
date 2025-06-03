package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
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
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname: %w", err)
		return
	}
	if !*command.help && !*command.version {
		ips, err := net.LookupIP(hostname)
		if err != nil {
			fmt.Println("Error getting hostip: %w", err)
			return
		}
		for _, ip := range ips {
			ipv4 := ip.To4()
			if ipv4 != nil {
				hostid := (uint32(ipv4[0]) << 24) | (uint32(ipv4[1]) << 16) | (uint32(ipv4[2]) << 8) | uint32(
					ipv4[3],
				)
				fmt.Printf("%08x\n", hostid)
				return
			}
		}
	}
	if !*command.help && *command.version {
		fmt.Println(version)
		return
	}
	if *command.help && !*command.version {
		fmt.Println(
			"\n\nghostid [OPTION]\n\n",
			"Description:\n",
			"\tPrint the numeric identifier (in hexadecimal) for the current host.\n\n",
			"--help display this help and exit\n\n",
			"--version output version information and exit",
		)
		return
	}
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		fmt.Println("Error converting hex to bin: %w", err)
		return "", err
	}
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}
