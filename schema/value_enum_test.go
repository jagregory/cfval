package schema

import (
	"testing"
)

func TestEnumValidation(t *testing.T) {
	template := &Template{}
	tr := NewTemplateResource(template)
	context := []string{}

	enum := EnumValue{[]string{"a", "b", "c"}}

	if _, errs := enum.Validate(Schema{}, "", tr, context); errs == nil {
		t.Error("Enum should fail on empty string")
	}

	if _, errs := enum.Validate(Schema{}, "d", tr, context); errs == nil {
		t.Error("Enum should fail on anything which isn't a valid option")
	}

	if _, errs := enum.Validate(Schema{}, "b", tr, context); errs != nil {
		t.Error("Enum should pass on a valid option")
	}
}
