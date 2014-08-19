package rules

import (
	"strings"
)

var (
	YamlRules []Rule = []Rule{
		header,
	}
)

func header(c []byte, r Reporter) {
	header := strings.SplitN(string(c), "\n", 2)[0]
	if header != "#cloud-config" {
		r.Error(1, "must begin with #cloud-config")
	}
}
