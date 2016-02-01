package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type TemplateResource struct {
	Template   *Template
	Definition Resource
	Properties map[string]interface{}
	Metadata   map[string]interface{}
}

func (tr TemplateResource) Validate(context []string) (bool, []reporting.Failure) {
	failures := make([]reporting.Failure, 0, 50)

	if ok, errs := tr.Definition.Validate(tr, context); !ok {
		failures = append(failures, errs...)
	}

	if ok, errs := Json.Validate(tr.Metadata, tr, append(context, "Metadata")); !ok {
		failures = append(failures, errs...)
	}

	return len(failures) == 0, failures
}

func (tr TemplateResource) HasProperty(name string, expected interface{}) bool {
	if value, found := tr.Properties[name]; found {
		return value == expected
	}

	return false
}

func NewUnrecognisedResource(template *Template, awsType string) TemplateResource {
	return TemplateResource{
		Template: template,
		Definition: Resource{
			ValidateFunc: func(tr TemplateResource, context []string) (bool, []reporting.Failure) {
				return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), context)}
			},
		},
	}
}

type Template struct {
	Resources  map[string]TemplateResource
	Parameters map[string]Parameter
}

func (t Template) Validate() (bool, []reporting.Failure) {
	errors := make([]reporting.Failure, 0, 100)

	for logicalId, resource := range t.Resources {
		if ok, errs := resource.Validate([]string{"Resources", logicalId}); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
