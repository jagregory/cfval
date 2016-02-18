package schema

import (
	"fmt"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Attributes   map[string]Schema
	Properties   Properties
	ReturnValue  Schema
	ValidateFunc func(constraints.CurrentResource, Context) (reporting.ValidateResult, reporting.Reports)
}

func (rd Resource) Validate(tr constraints.CurrentResource, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(tr, ctx)
	}

	failures, visited := rd.Properties.Validate(tr, ctx)

	// Reject any properties we weren't expecting
	for _, key := range tr.Properties() {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for %s", key, rd.AwsType), ctx.Push(key).Path))
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func (rd Resource) TargetType() PropertyType {
	return rd.ReturnValue.TargetType()
}

func (r Resource) PropertyDefault(name string) interface{} {
	return r.Properties[name].Default
}

func NewUnrecognisedResource(awsType string) Resource {
	return Resource{
		ValidateFunc: func(resource constraints.CurrentResource, ctx Context) (reporting.ValidateResult, reporting.Reports) {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), ctx.Path)}
		},
	}
}
