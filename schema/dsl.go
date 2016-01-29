package schema

import (
	"fmt"
	"strings"

	"github.com/jagregory/cfval/reporting"
)

func EnumSchema(options ...string) Schema {
	return Schema{
		Type: TypeEnum,
		ValidateFunc: func(value interface{}, t Template, context []string) (bool, []reporting.Failure) {
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
		},
	}
}

func ArrayOf(schema Schema) Schema {
	schema.Array = true
	return schema
}

func Required(schema Schema) Schema {
	schema.Required = true
	return schema
}
