package schema

import (
	"fmt"
	"strings"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

//go:generate stringer -type=ValueType

type ValueType int

const (
	ValueUnknown ValueType = iota
	ValueString
	ValueBool
	ValueNumber
)

func (from ValueType) Same(to PropertyType) bool {
	if vt, ok := to.(ValueType); ok {
		return vt == from
	}

	return false
}

func (t ValueType) IsArray() bool {
	return false
}

// TODO: This really feels like it can be simplified
func (from ValueType) CoercibleTo(to PropertyType) Coercion {
	if from == ValueUnknown || to == ValueUnknown {
		return CoercionBegrudgingly
	}

	if to.Same(JSON) {
		return CoercionAlways
	}

	if to == from {
		return CoercionAlways
	}

	if to == ValueString {
		return CoercionBegrudgingly
	}

	if from == ValueString {
		return CoercionBegrudgingly
	}

	if to == ValueBool || to == ValueNumber || from == ValueBool || from == ValueNumber {
		return CoercionNever
	}

	return CoercionBegrudgingly
}

func (vt ValueType) Describe() string {
	return strings.TrimPrefix(vt.String(), "Value")
}

func (ValueType) PropertyDefault(string) (interface{}, bool) {
	return nil, false
}

func (vt ValueType) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if isValidValueForType(vt, value) {
		return reporting.ValidateOK, nil
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctionsAll)
	case map[string]interface{}:
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Object used in %s property", vt.Describe())}
	}

	errs := coerce(valueTypeOf(value), vt, ctx)
	if errs != nil {
		// We've either dangerously coerced, or failed to coerce, the value here
		// so we should abort further validations to prevent any type-specific
		// validations from running. e.g. StringLength or IntegerRange
		return reporting.ValidateAbort, reporting.Safe(errs)
	}

	return reporting.ValidateOK, reporting.Safe(errs)
}

// TODO: is this needed anymore? Coercion might just catch this behaviour
func isValidValueForType(vt ValueType, value interface{}) bool {
	switch vt {
	case ValueBool:
		if _, ok := value.(bool); ok {
			return true
		}
	case ValueString:
		if _, ok := value.(string); ok {
			return true
		}
	case ValueNumber:
		if _, ok := value.(float64); ok {
			return true
		}
	}

	return false
}

func convert(v interface{}) string {
	switch t := v.(type) {
	case PropertyType:
		return t.Describe()
	case bool:
		return "Bool"
	case float64:
		return "Number"
	case string:
		return "String"
	case map[string]interface{}:
		return "Object"
	case []interface{}:
		return "List<?>"
	default:
		return "Unknown"
	}
}

func valueTypeOf(v interface{}) ValueType {
	switch t := v.(type) {
	case bool:
		return ValueBool
	case float64:
		return ValueNumber
	case string:
		return ValueString
	case int:
		panic("Numbers are parsed as float64 in JSON, if you're seeing this panic either things have changed in Go's JSON parsing or you're using an int in a test")
	default:
		panic(fmt.Errorf("Didn't expect %T here, should've been handled further up", t))
	}
}
