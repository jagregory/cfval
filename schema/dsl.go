package schema

import (
	"fmt"
	"regexp"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func SingleValueValidate(expected interface{}) ValidateFunc {
	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if value != expected {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Value must be %d but is %d", expected, value), context)}
		}

		return reporting.ValidateOK, nil
	}
}

func RegexpValidate(pattern, message string) ValidateFunc {
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, template, definitions, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		if re.MatchString(value.(string)) {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(message, context)}
	}
}

func IntegerRangeValidate(start, end float64) ValidateFunc {
	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueNumber.Validate(property, value, self, template, definitions, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		floatValue := value.(float64)

		if floatValue < start || floatValue > end {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Value must be between %f and %f", start, end), context)}
		}

		return reporting.ValidateOK, nil
	}
}

func StringLengthValidate(min, max int) ValidateFunc {
	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		if result, errs := ValueString.Validate(property, value, self, template, definitions, context); result == reporting.ValidateAbort || errs != nil {
			return reporting.ValidateOK, errs
		}

		str := value.(string)

		if len(str) < min || len(str) > max {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("String length must be between %d and %d", min, max), context)}
		}

		return reporting.ValidateOK, nil
	}
}

func NumberOptions(numbers ...float64) ValidateFunc {
	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		for _, n := range numbers {
			if n == value.(float64) {
				return reporting.ValidateOK, nil
			}
		}
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Number must be one of %v", numbers), context)}
	}
}

// TODO: this is really dumb, but it's late and I'm tired
func match(left []string, right []interface{}) bool {
	if len(left) != len(right) {
		return false
	}

	set := make(map[string]bool)

	for _, item := range left {
		set[item] = false
	}

	for _, item := range right {
		str := item.(string) // TODO: this will fail if the list contains a {ref} or something

		if _, found := set[str]; found {
			set[str] = true
		} else {
			return false
		}
	}

	for _, found := range set {
		if !found {
			return false
		}
	}

	return true
}

func contains(all []string, one string) bool {
	for _, item := range all {
		if item == one {
			return true
		}
	}

	return false
}

// TODO: fixme
func FixedArrayValidate(options ...[]string) ValidateFunc {
	return func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
		for _, option := range options {
			if match(option, value.([]interface{})) {
				return reporting.ValidateOK, nil
			}
		}

		// TODO: this should be []TypeString but we can't specify that with this method
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Invalid list value: %s, expected one of [%s]", value, options), context)}
	}
}
