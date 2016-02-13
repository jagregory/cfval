package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Attributes   map[string]Schema
	Properties   Properties
	ReturnValue  Schema
	ValidateFunc func(TemplateResource, []string) (reporting.ValidateResult, reporting.Failures)
}

func (rd Resource) Validate(tr TemplateResource, context []string) (reporting.ValidateResult, reporting.Failures) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(tr, context)
	}

	failures, visited := rd.Properties.Validate(tr, tr.Properties, context)

	// Reject any properties we weren't expecting
	for key := range tr.Properties {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for %s", key, rd.AwsType), append(context, key)))
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
