package schema

import (
	"fmt"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type GetAtt struct {
	source     Schema
	definition []interface{}
}

func NewGetAtt(source Schema, definition []interface{}) GetAtt {
	return GetAtt{source, definition}
}

func (ga GetAtt) Validate(template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.ValidateResult, reporting.Reports) {
	if len(ga.definition) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(ga.definition)), path)}
	}

	if resourceID, ok := ga.definition[0].(string); ok {
		if resource, ok := template.Resources[resourceID]; ok {
			if attributeName, ok := ga.definition[1].(string); ok {
				definition := definitions.Lookup(resource.Type)
				// TODO: BUG this line below should be attribute, ok
				if resource, ok := definition.Attributes[attributeName]; ok {
					// TODO: make this common, so GetAtt and others can use it
					targetType := resource.Type
					switch targetType.CoercibleTo(ga.source.Type) {
					case CoercionNever:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, attributeName, targetType.Describe(), ga.source.Type.Describe()), path)}
					case CoercionBegrudgingly:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(fmt.Sprintf("GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, attributeName, targetType.Describe(), ga.source.Type.Describe()), path)}
					}

					return reporting.ValidateAbort, nil
				}
			}

			// attribute not found on resource
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt %s.%s is not an attribute", resourceID, ga.definition[1]), path)}
		}

		// resource not found
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a resource", resourceID), path)}
	}

	// resource not a string
	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a valid resource name", ga.definition[0]), path)}
}
