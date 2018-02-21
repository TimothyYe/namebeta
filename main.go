package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	options := parseArgs(os.Args)
	if options == nil {
		displayUsage()
		os.Exit(0)
	}

	if err := query(options); err != nil {
		color.Red(fmt.Sprintf("%s %s \r\n", crossSymbol, err.Error()))
		os.Exit(1)
	}
}
