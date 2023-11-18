package main

import (
	"github.com/mingchen-ai-code/ExampleCode/validator/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// create a bson.M object
	adata := []bson.M{
		{"a": "a", "b": "b"},
		{"c": "a", "d": "b"},
	}

	// create a validator object
	validator := utils.CreateValidator()

	// add a rule to the validator
	rule := &utils.ChangeAtoB{}

	validator.Preview(rule, adata)
	rules := []utils.IRule{
		rule,
	}

	for _, rule := range rules {
		validator.AddRule(rule)
	}
	validator.Update(adata)

}
