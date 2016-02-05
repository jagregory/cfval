package schema

import "github.com/jagregory/cfval/reporting"

type Parameter struct {
	Type ValueType
}

func (Parameter) Validate([]string) (bool, []reporting.Failure) {
	return true, nil
}

func (p Parameter) TargetType() ValueType {
	return p.Type
}
