package utils

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type ChangeAtoB struct{}

func (c *ChangeAtoB) Correct(data bson.M) (RuleOutput, error) {
	// data := bson.M{"a": 1, "b": 2}
	// for all fields, change its value from 'a' to 'b'
	changed := false
	for k := range data {
		if val, ok := data[k].(string); ok {
			tmp := strings.ReplaceAll(val, "a", "b")
			if tmp == val {
				continue
			}
			changed = true
			data[k] = tmp
		}
	}

	return RuleOutput{Output: data, Change: changed}, nil
}

func (c *ChangeAtoB) GetName() string {
	return "ChangeAtoB"
}
