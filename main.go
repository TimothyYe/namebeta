package main

import (
	"os"
)

func main() {

	if len(os.Args) == 1 {
		displayUsage()
		os.Exit(0)
	}

	domain, withMore, withWhois := parseArgs(os.Args)
	query(domain, withMore, withWhois)
}
