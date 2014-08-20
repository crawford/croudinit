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
	for _, e := range report.Errors() {
		fmt.Printf("Error   <stdin>:%d %s\n", e.Line, e.Message)
	}
	for _, w := range report.Warnings() {
		fmt.Printf("Warning <stdin>:%d %s\n", w.Line, w.Message)
	}
}
