package requests_validators_rules

type Rule interface {
	validate(arg interface{}) bool
}
