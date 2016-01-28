package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateResourceProperty(r Resource, value interface{}, t Template, context []string) (bool, []Failure) {
	if properties, ok := value.(map[string]interface{}); ok {
		return r.Validate(t, properties, context)
	}

	return false, []Failure{NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, r.AwsType), context)}
}

func validateProperty(s Schema, value interface{}, t Template, context []string) (bool, []Failure) {
	if resource, ok := s.Type.(Resource); ok {
		return validateResourceProperty(resource, value, t, context)
	}

	if ok := validateValueType(s.Type, value, t, context); !ok {
		if complex, ok := value.(map[string]interface{}); ok {
			return validateBuiltinFns(complex, t, context)
		}

		return false, []Failure{NewInvalidTypeFailure(s.Type, value, context)}
	}

	if s.ValidateFunc != nil {
		return s.ValidateFunc(value, t, context)
	}

	return true, nil
}

func validateValueType(valueType interface{}, value interface{}, t Template, context []string) bool {
	switch valueType {
	case TypeBool:
		if _, ok := value.(bool); ok {
			return true
		}
	case TypeEnum:
		fallthrough
	case TypeString:
		if _, ok := value.(string); ok {
			return true
		}
	}

	return false
}

func validateRef(ref string, t Template, context []string) (bool, []Failure) {
	if _, ok := t.Resources[ref]; ok {
		// ref is to a resource and we've found it
		// TODO: validate resource ref value is correct type for property
		return true, nil
	} else if _, ok := t.Parameters[ref]; ok {
		// ref is to a parameter and we've found it
		// TODO: validate parameter type is correct for property
		return true, nil
	}

	return false, []Failure{NewFailure(fmt.Sprintf("Ref '%s' is not a resource or parameter", ref), context)}
}

func validateBuiltinFns(value map[string]interface{}, t Template, context []string) (bool, []Failure) {
	if ref, ok := value["Ref"]; ok {
		if refstr, ok := ref.(string); ok {
			return validateRef(refstr, t, context)
		}

		return false, []Failure{NewFailure(fmt.Sprintf("Ref has invalid value '%s'", ref), context)}
	}

	if _, ok := value["Fn::Find"]; ok {
		return false, []Failure{NewFailure("Value is an Fn::Find but this isn't supported yet", context)}
	}

	if _, ok := value["Fn::Join"]; ok {
		return false, []Failure{NewFailure("Value is an Fn::Join but this isn't supported yet", context)}
	}

	if _, ok := value["Fn::GetAtt"]; ok {
		return false, []Failure{NewFailure("Value is an Fn::GetAtt but this isn't supported yet", context)}
	}

	return false, []Failure{NewFailure("Value is a map but isn't a builtin", context)}
}

type ValidateFunc func(interface{}, Template, []string) (bool, []Failure)

type Schema struct {
	Array        bool
	Required     bool
	Type         interface{}
	ValidateFunc ValidateFunc
}

func (s Schema) Validate(value interface{}, t Template, context []string) (bool, []Failure) {
	if !s.Required && value == nil {
		return true, nil
	}

	failures := make([]Failure, 0, 20)

	if s.Array {
		for i, item := range value.([]interface{}) {
			if ok, errs := validateProperty(s, item, t, append(context, strconv.Itoa(i))); !ok {
				failures = append(failures, errs...)
			}
		}
	} else {
		if ok, errs := validateProperty(s, value, t, context); !ok {
			failures = append(failures, errs...)
		}
	}

	return len(failures) == 0, failures
}

//go:generate stringer -type=ValueType

type ValueType int

const (
	TypeEnum ValueType = iota
	TypeString
	TypeBool
	TypeInteger
)

func EnumSchema(options ...string) Schema {
	return Schema{
		Type: TypeEnum,
		ValidateFunc: func(value interface{}, t Template, context []string) (bool, []Failure) {
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
				} else {
					return false, []Failure{NewFailure(fmt.Sprintf("Invalid enum option %s, expected one of [%s]", str, strings.Join(options, ", ")), context)}
				}
			}

			return false, []Failure{NewInvalidTypeFailure(TypeEnum, value, context)}
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
