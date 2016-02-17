package schema

import (
	"fmt"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Attributes   map[string]Schema
	Properties   Properties
	ReturnValue  Schema
	ValidateFunc func(parse.TemplateResource, []string) (reporting.ValidateResult, reporting.Reports)
}

func (rd Resource) Validate(tr parse.TemplateResource, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(tr, context)
	}

	failures, visited := rd.Properties.Validate(tr, definitions, tr.Properties, context)

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

func NewUnrecognisedResource(awsType string) Resource {
	return Resource{
		ValidateFunc: func(tr parse.TemplateResource, context []string) (reporting.ValidateResult, reporting.Reports) {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), context)}
		},
	}
}
