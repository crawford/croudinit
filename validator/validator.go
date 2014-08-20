package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/crawford/crowdconfig/validator/report"
	"github.com/crawford/crowdconfig/validator/rules"
)

func main() {
	config, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	report := &report.Report{}
	for _, r := range rules.Rules {
		r(config, report)
	}
	for _, e := range report.Entries() {
		if e.IsError() {
			fmt.Printf("Error   <stdin>:%s\n", e)
		} else if e.IsWarning() {
			fmt.Printf("Warning <stdin>:%s\n", e)
		}
	}
}
