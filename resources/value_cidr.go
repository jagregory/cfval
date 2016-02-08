package resources

import (
	"fmt"
	"regexp"

	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

// TODO: Can we switch this to the regexp validate?
const cidrPattern = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`

var cidr FuncType = func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	if ok, _ := regexp.MatchString(cidrPattern, value.(string)); !ok {
		return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Cidr %s is invalid", value), context)}
	}

	return reporting.ValidateOK, nil
}
