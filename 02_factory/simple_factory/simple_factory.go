package simple_factory

import "fmt"

type IRuleConfigParser interface {
	Parse()
}

type jsonRuleConfigParser struct{}

func (j *jsonRuleConfigParser) Parse() {
	fmt.Println("json parse")
}

type yamlRuleConfigParser struct{}

func (y *yamlRuleConfigParser) Parse() {
	fmt.Println("yaml parse")
}

func NewRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return &jsonRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	}
	return nil
}
