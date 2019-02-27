package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

const (
	domainURL = "https://namebeta.com/api/query"
	whoisURL  = "https://namebeta.com/api/whois"
	referURL  = "https://namebeta.com/search/%s"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36"
)

func query(o *options) error {
	if o.Whois {
		return queryWhois(o.Domain)
	}

	return queryDomain(o.Domain, o.WithMore)
}

func queryWhois(domain string) error {

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	if err := s.Color("green"); err != nil {
		color.Red("Cannot set color")
		os.Exit(1)
	}
	s.Start()

	params := map[string]string{"domain": domain}
	result, err := getDomainInfo(whoisURL, domain, params)
	s.Stop()

	if err != nil {
		return err
	}

	if result[0].(bool) {
		fmt.Println()

		status := result[1].(map[string]interface{})["status"].(float64)

		if status == 1 {
			color.Red("NOT FOUND.")
		} else {
			color.Cyan(result[1].(map[string]interface{})["whois"].(string))
		}

		return nil
	}

	return fmt.Errorf("Failed to query domain: %s", domain)
}

func queryDomain(domain string, withMore bool) error {
	params := map[string]string{"q": domain}

	var resultMore []interface{}

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	if err := s.Color("green"); err != nil {
		color.Red("Failed to set color")
		os.Exit(1)
	}
	s.Start()

	result, err := getDomainInfo(domainURL, domain, params)

	// For more option, query special domains
	if withMore {
		params["special"] = "1"
		resultMore, err = getDomainInfo(domainURL, domain, params)

		if len(resultMore) > 0 && resultMore[0].(bool) {
			result[2] = append(result[2].([]interface{}), resultMore[2].([]interface{})...)
		}
	}

	s.Stop()

	if err != nil {
		return err
	}

	if result[0].(bool) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain", "Available"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, v := range result[2].([]interface{}) {
			data := v.([]interface{})
			row := []string{data[0].(string)}

			switch data[1].(float64) {
			case 1:
				if runtime.GOOS == "windows" {
					row = append(row, checkSymbol)
				} else {
					row = append(row, color.GreenString(checkSymbol))
				}
			case 2:
				if runtime.GOOS == "windows" {
					row = append(row, crossSymbol)
				} else {
					row = append(row, color.GreenString(crossSymbol))
				}
			}

			table.Append(row)
		}
		table.Render()

		return nil
	}

	return fmt.Errorf("Failed to query domain: %s", domain)
}
