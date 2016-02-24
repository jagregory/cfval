package schema

import (
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

func (res NestedResource) Same(to PropertyType) bool {
	if nr, ok := to.(NestedResource); ok {
		return nr.Description == res.Description
	}

	return false
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
		tnr := parse.NewTemplateResource(property.Type.Describe(), values)

		nestedResourceContext := NewResourceContext(ctx, ResourceWithDefinition{tnr, property.Type})
		failures, visited := res.Properties.Validate(nestedResourceContext)

		// Reject any properties we weren't expecting
		for key := range res.Properties {
			if !visited[key] {
				failures = append(failures, reporting.NewFailure(PropertyContextAdd(ctx, key), "Unknown property '%s' for nested %s", key, res.Description))
			}
		}

		if len(failures) == 0 {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, failures
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid type %T for nested resource %s", value, res.Description)}
}
