package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validateType(valueType interface{}, value interface{}, t Template, context []string) bool {
	if _, ok := value.(bool); valueType == TypeBool && ok {
		return true
	} else if _, ok := value.(string); (valueType == TypeString || valueType == TypeEnum) && ok {
		return true
	}

	return false
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

	var validate = func(item interface{}, ctx []string) (bool, []Failure) {
		basicValidityPassed := validateType(s.Type, item, t, ctx)

		if basicValidityPassed && s.ValidateFunc != nil {
			return s.ValidateFunc(item, t, ctx)
		} else if !basicValidityPassed {
			return false, []Failure{NewInvalidTypeFailure(s.Type, item, ctx)}
		}

		return true, nil
	}

	failures := make([]Failure, 0, 20)

	if s.Array {
		for i, item := range value.([]interface{}) {
			if ok, errs := validate(item, append(context, strconv.Itoa(i))); !ok {
				failures = append(failures, errs...)
			}
		}
	} else {
		if ok, errs := validate(value, context); !ok {
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
