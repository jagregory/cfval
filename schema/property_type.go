package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type ArrayPropertyType struct {
	PropertyType
}

func (apt ArrayPropertyType) Unwrap() PropertyType {
	return apt.PropertyType
}

func (pt ArrayPropertyType) Describe() string {
	return fmt.Sprintf("List<%s>", pt.PropertyType.Describe())
}

func (pt ArrayPropertyType) Same(to PropertyType) bool {
	if apt, ok := to.(ArrayPropertyType); ok {
		return pt.PropertyType.Same(apt.PropertyType)
	}

	return false
}

func (pt ArrayPropertyType) CoercibleTo(to PropertyType) Coercion {
	if pt.Same(to) {
		return CoercionAlways
	} else if vt, ok := to.(ValueType); ok && vt == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func Multiple(pt PropertyType) PropertyType {
	if _, ok := pt.(ArrayPropertyType); ok {
		panic("Multiple(Multiple(...)) call detected")
	}

	return ArrayPropertyType{pt}
}

type PropertyType interface {
	// Describe returns a human-readable description of the type in AWS, which
	// could be the AWS Resource Type or just any arbitrary description.
	Describe() string

	// Same returns true when two PropertyTypes represent the same AWS type.
	Same(PropertyType) bool

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
