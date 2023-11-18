package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type RuleOutput struct {
	Output bson.M
	Change bool
}

type IRule interface {
	Correct(bson.M) (RuleOutput, error)
	GetName() string
}

type Validator struct {
	rules []IRule
}

func CreateValidator() *Validator {
	return &Validator{}
}

func (v *Validator) AddRule(rule IRule) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Preview(rule IRule, data []bson.M) {
	for _, d := range data {
		ruleOutput, err := rule.Correct(d)
		if err == nil && ruleOutput.Change {
			fmt.Println("Before: ", rule.GetName(), d)
			fmt.Println("After: ", rule.GetName(), ruleOutput.Output)
		}
	}

}
func (v *Validator) Update(data []bson.M) {
	for _, rule := range v.rules {
		correctOutputByRule := make([]bson.M, 0)
		for _, d := range data {
			ruleOutput, err := rule.Correct(d)
			if err == nil && ruleOutput.Change {
				correctOutputByRule = append(correctOutputByRule, ruleOutput.Output)
			}
		}
		fmt.Println("CorrectOutputByRule: ", rule.GetName(), correctOutputByRule)
	}
}
