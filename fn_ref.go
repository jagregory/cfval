package main

import "fmt"

type Ref struct {
	Target string
}

func (ref Ref) Validate(t Template, context []string) (bool, []Failure) {
	if _, ok := t.Parameters[ref.Target]; !ok {
		return false, []Failure{NewFailure(fmt.Sprintf("Invalid Ref '%s'", ref.Target), context)}
	}
	return true, nil
}
