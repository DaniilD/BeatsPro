package requests_validators_rules

type Rule interface {
	Validate(arg interface{}) bool
}
