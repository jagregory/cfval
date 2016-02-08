package schema

import "testing"

func TestParameterTargetType(t *testing.T) {
	if (Parameter{Schema: Schema{Type: ValueNumber}}).TargetType() != ValueNumber {
		t.Error("Parameter TargetType should match Ref type")
	}

	if (Parameter{}).TargetType() != ValueUnknown {
		t.Error("Parameter without Ref should return TypeUnknown for TargetType")
	}
}
