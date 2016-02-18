package schema

import (
	"fmt"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type NestedResource struct {
	Description string
	Properties
}

func (res NestedResource) Describe() string {
	return res.Description
}

func (NestedResource) CoercibleTo(PropertyType) Coercion {
	return CoercionNever
}

// TODO: This is all a bit hairy. We shouldn't need to be creating the
// 			 TemplateNestedResource here, ideally `self` should already refer to
//			 one and value should already be a map[string]inteface{}
func (res NestedResource) Validate(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.ValidateResult, reporting.Reports) {
	if values, ok := value.(map[string]interface{}); ok {
		tnr := parse.NewTemplateResource(template, property.Type.Describe(), values)
		failures, visited := res.Properties.Validate(ResourceWithDefinition{tnr, property.Type}, template, definitions, path)

		// Reject any properties we weren't expecting
		for key := range res.Properties {
			if !visited[key] {
				failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for nested %s", key, res.Description), append(path, key)))
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, res.Description), path)}
}
