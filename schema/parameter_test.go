package schema

import "testing"

func TestParameterTargetType(t *testing.T) {
	if (Parameter{Type: TypeInteger}).TargetType() != TypeInteger {
		t.Error("Parameter TargetType should match Ref type")
	}

	if (Parameter{}).TargetType() != TypeUnknown {
		t.Error("Parameter without Ref should return TypeUnknown for TargetType")
	}
}
