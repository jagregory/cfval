package schema

import "github.com/jagregory/cfval/parse"

type PropertyDefaults interface {
	PropertyDefault(name string) interface{}
}

type ResourceWithDefinition struct {
	parse.TemplateResource
	PropertyDefaults
}
