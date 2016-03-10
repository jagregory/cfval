package schema

import "fmt"

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
	} else if to.Same(JSON) {
		singleItemCoercion := pt.Unwrap().CoercibleTo(JSON)
		return singleItemCoercion
	} else if at, ok := to.(ArrayPropertyType); ok {
		singleItemCoercion := pt.Unwrap().CoercibleTo(at.Unwrap())
		return singleItemCoercion
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
