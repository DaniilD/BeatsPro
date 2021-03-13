package requests_validators_rules

import (
	"reflect"
)

type NotEmptyRule struct {
}

func NewNotEmptyRule() *NotEmptyRule {
	return &NotEmptyRule{}
}

func (notEmptyRule *NotEmptyRule) Validate(arg interface{}) bool {
	if reflect.TypeOf(arg).Kind() == reflect.String {
		if arg == "" {
			return false
		}
	} else {
		if reflect.ValueOf(arg).IsNil() {
			return false
		}
	}

	return true
}
