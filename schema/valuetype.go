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
		case parse.Ref:
			return validateRef(t, ctx)
		case parse.FindInMap:
			return validateFindInMap(t, PropertyContextAdd(ctx, "Fn::FindInMap"))
		case parse.Join:
			return validateJoin(t, PropertyContextAdd(ctx, "Fn::Join"))
		case parse.GetAtt:
			return validateGetAtt(t, PropertyContextAdd(ctx, "Fn::GetAtt"))
		case map[string]interface{}:
			builtinResult, errs := ValidateBuiltinFns(t, ctx)
			if errs != nil {
				return reporting.ValidateOK, errs
			}

			if builtinResult == reporting.ValidateAbort {
				return reporting.ValidateAbort, nil
			}

			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Value is a map but isn't a builtin")}
		}

		return reporting.ValidateOK, reporting.Reports{reporting.NewInvalidTypeFailure(ctx, vt, convert(value))}
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

func convert(v interface{}) ValueType {
	switch v.(type) {
	case bool:
		return ValueBool
	case float64:
		return ValueNumber
	case string:
		return ValueString
	default:
		return ValueUnknown
	}
}
