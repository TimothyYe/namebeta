package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

var (
	Version = "0.1"
	logo    = `
                             _                         
                            | |            _           
 ____    ____  ____    ____ | | _    ____ | |_    ____ 
|  _ \  / _  ||    \  / _  )| || \  / _  )|  _)  / _  |
| | | |( ( | || | | |( (/ / | |_) )( (/ / | |__ ( ( | |
|_| |_| \_||_||_|_|_| \____)|____/  \____) \___) \_||_|

NameBeta V%s

Insired by https://namebeta.com
https://github.com/Timothy/namebeta

`
)

const (
	more         = "-m"
	whois        = "-w"
	help         = "-h"
	checkSymbol  = "\u2714"
	crossSymbol  = "\u2716"
	circleSymbol = "\u25CF"
)

func displayUsage() {
	color.Cyan(logo, Version)
	color.Cyan("Usage:")
	color.Cyan("namebeta <domain to query>    Query with input domain")
	color.Cyan("namebeta -m <domain to query> Query more results with input domain")
	color.Cyan("namebeta -w <domain to query> Query WHOIS infomation with input domain")
	color.Cyan("namebeta -h                   Display usage and help")
}

func parseArgs(args []string) (string, bool, bool) {

	switch args[1] {
	case help:
		displayUsage()
		os.Exit(0)
	case more:
		if len(args) == 2 {
			displayUsage()
			os.Exit(1)
		}
		return args[2], true, false
	case whois:
		if len(args) == 2 {
			displayUsage()
			os.Exit(1)
		}
		return args[2], false, true
	default:
		return args[1], false, false
	}

	return "", true, true
}

func whoisQuery(domain string) []interface{} {
	var result []interface{}
	param := map[string]string{}
	param["domain"] = domain

	request := gorequest.New()
	_, body, _ := request.Post(whoisURL).
		Type("form").
		Set("User-Agent", userAgent).
		Set("Refer", fmt.Sprintf(referURL, domain)).
		SendMap(param).End()

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		color.Red(fmt.Sprintf("%s Failed to query WHOIS information for domain: %s \r\n", crossSymbol, domain))
		os.Exit(1)
	}

	return result
}

func domainQuery(domain string, param map[string]string) []interface{} {
	var result []interface{}

	request := gorequest.New()
	_, body, _ := request.Post(domainURL).
		Type("form").
		Set("User-Agent", userAgent).
		Set("Refer", fmt.Sprintf(referURL, domain)).
		SendMap(param).End()

	if err := json.Unmarshal([]byte(body), &result); err != nil {
		color.Red(fmt.Sprintf("%s Failed to query domain: %s \r\n", crossSymbol, domain))
		os.Exit(1)
	}

	return result
}
