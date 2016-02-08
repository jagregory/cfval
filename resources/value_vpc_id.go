package resources

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var vpcID FuncType = func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	// TODO: VpcId
	return reporting.ValidateOK, nil
}
