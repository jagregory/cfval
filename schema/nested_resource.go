package schema

import (
	"fmt"

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
func (res NestedResource) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if values, ok := value.(map[string]interface{}); ok {
		property := ctx.Property()
		tnr := parse.NewTemplateResource(ctx.Template(), property.Type.Describe(), values)

		nestedResourceContext := NewResourceContext(ctx, ResourceWithDefinition{tnr, property.Type})
		failures, visited := res.Properties.Validate(nestedResourceContext)

		// Reject any properties we weren't expecting
		for key := range res.Properties {
			if !visited[key] {
				failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for nested %s", key, res.Description), PropertyContextAdd(ctx, key).Path()))
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, res.Description), ctx.Path())}
}
