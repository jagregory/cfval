package schema

import (
	"fmt"
	"strings"

	"github.com/jagregory/cfval/reporting"
)

type EnumValue struct {
	Description string
	Options     []string
}

func (enum EnumValue) Describe() string {
	return enum.Description
}

func (EnumValue) PropertyDefault(string) interface{} {
	return nil
}

func (from EnumValue) CoercibleTo(to PropertyType) Coercion {
	if to == ValueString {
		return CoercionAlways
	} else if ft, ok := to.(EnumValue); ok && ft.Description == from.Description {
		return CoercionAlways
	} else if to == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func (enum EnumValue) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
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

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Invalid enum option %s, expected one of [%s]", value, strings.Join(enum.Options, ", ")), ctx)}
}
