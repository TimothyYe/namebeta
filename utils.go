package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/fatih/color"
)

var (
	// Version of namebeta
	Version = "0.1"
	logo    = `
                             _                         
                            | |            _           
 ____    ____  ____    ____ | | _    ____ | |_    ____ 
|  _ \  / _  ||    \  / _  )| || \  / _  )|  _)  / _  |
| | | |( ( | || | | |( (/ / | |_) )( (/ / | |__ ( ( | |
|_| |_| \_||_||_|_|_| \____)|____/  \____) \___) \_||_|

NameBeta V%s

Inspired by https://namebeta.com
https://github.com/TimothyYe/namebeta

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
	color.Cyan("namebeta -w <domain to query> Query WHOIS information with input domain")
	color.Cyan("namebeta -h                   Display usage and help")
}

type options struct {
	Domain   string
	WithMore bool
	Whois    bool
}

func parseArgs(args []string) *options {
	if len(os.Args) == 1 {
		return nil
	}

	switch args[1] {
	case more:
		if len(args) > 2 {
			return &options{args[2], true, false}
		}
	case whois:
		if len(args) > 2 {
			return &options{args[2], false, true}
		}
	default:
		return &options{args[1], false, false}
	}

	return nil
}

func getDomainInfo(endpoint string, domain string, params map[string]string) ([]interface{}, error) {
	var result []interface{}

	form := url.Values{}
	for k, v := range params {
		form.Add(k, v)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Referer", fmt.Sprintf(referURL, domain))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Failed to query domain: %s, endpoint: %s, error: %s", domain, endpoint, err.Error())
	}

	return result, nil
}
