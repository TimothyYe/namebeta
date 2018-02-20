package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	cli := parseArgs(os.Args)
	if cli == nil {
		displayUsage()
		os.Exit(0)
	}

	if err := query(cli); err != nil {
		color.Red(fmt.Sprintf("%s %s \r\n", crossSymbol, err.Error()))
		os.Exit(1)
	}
}
