package schema

import "github.com/jagregory/cfval/reporting"

type Parameter struct {
	Type string
}

func (Parameter) Validate([]string) (bool, []reporting.Failure) {
	return true, nil
}
