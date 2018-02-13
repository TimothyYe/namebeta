package main

import (
	"fmt"
	"os"
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

func query(domain string, withMore, withWhois bool) {
	if withWhois {
		queryWhois(domain)
	} else {
		queryDomain(domain, withMore)
	}
}

func queryWhois(domain string) {

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	s.Color("green")
	s.Start()

	result := whoisQuery(domain)
	s.Stop()

	if result[0].(bool) {
		fmt.Println()

		status := result[1].(map[string]interface{})["status"].(float64)

		if status == 1 {
			color.Red("NOT FOUND.")
		} else {
			color.Cyan(result[1].(map[string]interface{})["whois"].(string))
		}
	} else {
		color.Red(fmt.Sprintf("%s Failed to query domain: %s \r\n", crossSymbol, domain))
		os.Exit(1)
	}
}

func queryDomain(domain string, withMore bool) {
	param := map[string]string{}
	param["q"] = domain

	var resultMore []interface{}

	//Init spinner
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Prefix = "Querying... "
	s.Color("green")
	s.Start()

	result := domainQuery(domain, param)

	// For more option, query special domains
	if withMore {
		param["special"] = "1"
		resultMore = domainQuery(domain, param)
	}

	if len(resultMore) > 0 && resultMore[0].(bool) {
		result[2] = append(result[2].([]interface{}), resultMore[2].([]interface{})...)
	}

	s.Stop()

	if result[0].(bool) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Domain", "Available"})
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		for _, v := range result[2].([]interface{}) {
			data := v.([]interface{})
			row := []string{}
			row = append(row, data[0].(string))

			switch data[1].(float64) {
			case 1:
				row = append(row, color.GreenString(checkSymbol))
			case 2:
				row = append(row, color.RedString(crossSymbol))
			}

			table.Append(row)
		}
		table.Render()

	} else {
		color.Red(fmt.Sprintf("%s Failed to query domain: %s \r\n", crossSymbol, domain))
		os.Exit(1)
	}
}
