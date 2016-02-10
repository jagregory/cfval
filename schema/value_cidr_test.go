package schema

import "testing"

func TestCidrValidation(t *testing.T) {
	template := &Template{}
	tr := NewTemplateResource(template)
	context := []string{}

	if _, errs := CIDR.Validate(Schema{}, "", tr, context); errs == nil {
		t.Error("Cidr should fail on empty string")
	}

	if _, errs := CIDR.Validate(Schema{}, "abc", tr, context); errs == nil {
		t.Error("Cidr should fail on anything which isn't a cidr")
	}

	if _, errs := CIDR.Validate(Schema{}, "0.0.0.0/100", tr, context); errs == nil {
		t.Error("Cidr should fail on an invalid mask")
	}

	if _, errs := CIDR.Validate(Schema{}, "10.200.300.10/24", tr, context); errs == nil {
		t.Error("Cidr should fail on an invalid IP")
	}

	if _, errs := CIDR.Validate(Schema{}, "10.200.30.10/24", tr, context); errs != nil {
		t.Error("Cidr should pass with a valid cidr")
	}
}
