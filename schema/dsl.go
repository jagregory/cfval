package schema

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jagregory/cfval/reporting"
)

func EnumValidate(options ...string) ValidateFunc {
	return func(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
		if str, ok := value.(string); ok {
			found := false
			for _, option := range options {
				if option == str {
					found = true
					break
				}
			}

			if found {
				return true, nil
			}

			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Invalid enum option %s, expected one of [%s]", str, strings.Join(options, ", ")), context)}
		}

		return false, []reporting.Failure{reporting.NewInvalidTypeFailure(TypeEnum, value, context)}
	}
}

func RegexpValidate(pattern, message string) ValidateFunc {
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	return func(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
		if re.MatchString(value.(string)) {
			return true, nil
		}

		return false, []reporting.Failure{reporting.NewFailure(message, context)}
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

func FixedArrayValidate(options ...[]string) ArrayValidateFunc {
	return func(value []interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
		for _, option := range options {
			if match(option, value) {
				return true, nil
			}
		}

		// TODO: this should be []TypeString but we can't specify that with this method
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Invalid list value: %s, expected one of [%s]", value, options), context)}
	}
}

func Required(schema Schema) Schema {
	schema.Required = true
	return schema
}
