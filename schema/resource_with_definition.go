package schema

import "github.com/jagregory/cfval/parse"

type PropertyDefaults interface {
	PropertyDefault(name string) (interface{}, bool)
}

type ResourceWithDefinition struct {
	parse.TemplateResource
	PropertyDefaults
}

func (r ResourceWithDefinition) AwsType() string {
	return r.TemplateResource.Type
}

func (r ResourceWithDefinition) Id() string {
	return r.TemplateResource.LogicalID
}

func (r ResourceWithDefinition) PropertyValueOrDefault(name string) (interface{}, bool) {
	if v, ok := r.PropertyValue(name); ok {
		return v, true
	}

	if v, ok := r.PropertyDefault(name); ok {
		return v, true
	}

	return nil, false
}
