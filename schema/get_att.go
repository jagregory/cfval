package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type GetAtt struct {
	definition []interface{}
}

func NewGetAtt(definition []interface{}) GetAtt {
	return GetAtt{definition}
}

func (ga GetAtt) Validate(template *Template, context []string) (reporting.ValidateResult, reporting.Failures) {
	if len(ga.definition) != 2 {
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(ga.definition)), context)}
	}

	if resourceID, ok := ga.definition[0].(string); ok {
		if _, ok := template.Resources[resourceID]; ok {
			if _, ok := ga.definition[1].(string); ok {
				// TODO: Check attr is actually a valid attribute for the resource type
				return reporting.ValidateAbort, nil
			}
		}
		// resource not found
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a resource", resourceID), context)}
	}

	// resource not a string
	return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a valid resource name", ga.definition[0]), context)}
}
