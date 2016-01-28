package main

import "fmt"

type TemplateResource struct {
	Definition Resource
	Properties map[string]interface{}
}

func NewUnrecognisedResource(awsType string) TemplateResource {
	return TemplateResource{
		Definition: Resource{
			ValidateFunc: func(t Template, properties map[string]interface{}, context []string) (bool, []Failure) {
				return false, []Failure{NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), context)}
			},
		},
	}
}

type Template struct {
	Resources  map[string]TemplateResource
	Parameters map[string]Parameter
}

func (t Template) Validate() (bool, []Failure) {
	errors := make([]Failure, 0, 100)

	for logicalId, resource := range t.Resources {
		if ok, errs := resource.Definition.Validate(t, resource.Properties, []string{"Resources", logicalId}); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
