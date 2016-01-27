package main

import "fmt"

type AvailabilityZone struct {
	StringOrBuiltinFns
}

func (az AvailabilityZone) Validate(t Template, context []string) (bool,[]Failure) {
	if ok, errs := az.StringOrBuiltinFns.Validate(t, context); !ok {
		return ok, errs
	}

	// TODO
	if az.StringOrBuiltinFns.Value != nil && *az.StringOrBuiltinFns.Value != "ap-southeast-2b" {
		return false, []Failure{NewFailure(fmt.Sprintf("Unrecognised AvailabilityZone %s", *az.StringOrBuiltinFns.Value), context)}
	}

	return true, nil
}
