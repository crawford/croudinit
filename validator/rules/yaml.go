package rules

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/crawford/crowdconfig/third_party/launchpad.net/goyaml"
)

var (
	YamlRules []Rule = []Rule{
		header,
		syntax,
	}
	goyamlError = regexp.MustCompile(`^YAML error: line (?P<line>[[:digit:]]+): (?P<msg>.*)$`)
)

func header(c []byte, r Reporter) {
	header := strings.SplitN(string(c), "\n", 2)[0]
	if header != "#cloud-config" {
		r.Error(1, "must begin with #cloud-config")
	}
}

func syntax(c []byte, r Reporter) {
	if err := goyaml.Unmarshal(c, &struct{}{}); err != nil {
		matches := goyamlError.FindStringSubmatch(err.Error())
		if l, err := strconv.Atoi(matches[1]); err == nil {
			m := matches[2]
			r.Error(l, m)
		} else {
			panic(err)
		}
	}
}
