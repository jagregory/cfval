package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type arrayPropertyType struct {
	PropertyType
}

func (arrayPropertyType) IsArray() bool {
	return true
}

func (pt arrayPropertyType) Describe() string {
	return fmt.Sprintf("List<%s>", pt.PropertyType.Describe())
}

func (pt arrayPropertyType) Same(to PropertyType) bool {
	if apt, ok := to.(arrayPropertyType); ok {
		return pt.PropertyType.Same(apt.PropertyType)
	}

	return false
}

func (pt arrayPropertyType) CoercibleTo(to PropertyType) Coercion {
	if pt.Same(to) {
		return CoercionAlways
	} else if vt, ok := to.(ValueType); ok && vt == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func Multiple(pt PropertyType) PropertyType {
	return arrayPropertyType{pt}
}

type PropertyType interface {
	// Describe returns a human-readable description of the type in AWS, which
	// could be the AWS Resource Type or just any arbitrary description.
	Describe() string

	// Same returns true when two PropertyTypes represent the same AWS type.
	Same(PropertyType) bool

	// IsArray returns true when the PropertyType represents an array of
	// another PropertyType
	IsArray() bool

	// Validate checks that the property is valid, including any built-in function
	// calls and stuff within the property.
	Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports)

	// CoercibleTo will return true for types which the value of this property can
	// be coerced into. e.g. A number can be coerced to a string
	// CoercionAlways means a type is always coercible to another
	// 	 e.g. all numbers are valid strings
	// CoercionNever means a type is never coercible to another
	//   e.g. a number is never a valid bool
	// CoercionBegrudgingly means a type can be coerced but results may vary
	//   e.g. a string can be coerced to a number, but only if it is numerically
	//        valid.
	//
	// CoerceAlways and CoercionBegrudgingly are equivalent right now, but in
	// future a warning may be issued for begrudging conversions.
	CoercibleTo(PropertyType) Coercion

	// PropertyDefault returns the default value for a property, if one is set.
	PropertyDefault(name string) (interface{}, bool)
}
