package schema

import (
	"fmt"
	"strings"

	"github.com/jagregory/cfval/reporting"
)

type EnumValue struct {
	Options []string
}

func (enum EnumValue) Validate(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, []reporting.Failure) {
	if result, errs := ValueString.Validate(property, value, self, context); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	if str, ok := value.(string); ok {
		found := false
		for _, option := range enum.Options {
			if option == str {
				found = true
				break
			}
		}

		if found {
			return reporting.ValidateOK, nil
		}
	}

	return reporting.ValidateOK, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Invalid enum option %s, expected one of [%s]", value, strings.Join(enum.Options, ", ")), context)}
}