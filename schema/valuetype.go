package schema

import (
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

	if ft, ok := to.(FuncType); ok && ft.Description == "JSON" {
		return CoercionNever
	}

	if to == from {
		return CoercionAlways
	}

	if to == ValueString {
		return CoercionAlways
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
	if ok := vt.validateValue(value); !ok {
		switch t := value.(type) {
		case parse.IntrinsicFunction:
			return ValidateIntrinsicFunctions(t, ctx, SupportedFunctionsAll)
		default:
			return reporting.ValidateOK, reporting.Reports{reporting.NewInvalidTypeFailure(ctx, convert(value), vt)}
		}
	}

	return reporting.ValidateOK, nil
}

func (vt ValueType) validateValue(value interface{}) bool {
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
