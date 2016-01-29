package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type TemplateResource struct {
	Definition Resource
	Properties map[string]interface{}
}

func (tr TemplateResource) HasProperty(name string, expected interface{}) bool {
	if value, found := tr.Properties[name]; found {
		return value == expected
	}

	return false
}

func NewUnrecognisedResource(awsType string) TemplateResource {
	return TemplateResource{
		Definition: Resource{
			ValidateFunc: func(t Template, properties map[string]interface{}, context []string) (bool, []reporting.Failure) {
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
		if ok, errs := resource.Definition.Validate(t, resource, resource.Properties, []string{"Resources", logicalId}); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
