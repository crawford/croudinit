package rules

type Reporter interface {
	Error(line int, message string)
	Warning(line int, message string)
}

type Rule func(contents []byte, reporter Reporter)

var (
	Rules []Rule = YamlRules
)
