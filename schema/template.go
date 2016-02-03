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
	Outputs    map[string]Output
}

func (t *Template) Validate() (bool, []reporting.Failure) {
	errors := make([]reporting.Failure, 0, 100)

	for logicalID, resource := range t.Resources {
		if ok, errs := resource.Validate([]string{"Resources", logicalID}); !ok {
			errors = append(errors, errs...)
		}
	}

	for parameterID, parameter := range t.Parameters {
		if ok, errs := parameter.Validate([]string{"Parameters", parameterID}); !ok {
			errors = append(errors, errs...)
		}
	}

	for outputID, output := range t.Outputs {
		if ok, errs := output.Validate(t, []string{"Outputs", outputID}); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
