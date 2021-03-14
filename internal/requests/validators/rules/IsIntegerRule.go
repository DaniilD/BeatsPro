package requests_validators_rules

import (
	"reflect"
	"strconv"
)

type IsIntegerRule struct {
}

func NewIsIntegerRule() *IsIntegerRule {
	return &IsIntegerRule{}
}

func (isIntegerRule *IsIntegerRule) Validate(arg interface{}) bool {
	if reflect.TypeOf(arg).Kind() == reflect.String {
		value := reflect.ValueOf(arg).String()
		if _, err := strconv.Atoi(value); err != nil {
			return false
		}
	} else {
		return false
	}

	return true
}
