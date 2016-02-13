package schema

import "testing"

func TestGetAtt(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{},
		},
	}
	context := []string{}

	if _, errs := NewGetAtt(nil).Validate(template, context); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"a", "b", "c"}).Validate(template, context); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"a"}).Validate(template, context); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"UnknownResource", "prop"}).Validate(template, context); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	// TODO: this test will eventually fail when we implement GetAtt prop validation
	// uncomment the following tests later
	if _, errs := NewGetAtt([]interface{}{"MyResource", "prop"}).Validate(template, context); errs != nil {
		t.Error("Should pass when valid resource used", errs)
	}

	// if _, errs := NewGetAtt([]interface{}{"MyResource", "BadProp"}).Validate(template, context); errs == nil {
	// 	t.Error("Should fail when invalid property used for type of resource", errs)
	// }
	//
	// if _, errs := NewGetAtt([]interface{}{"MyResource", "InstanceId"}).Validate(template, context); errs != nil {
	// 	t.Error("Should pass when valid property used for type of resource", errs)
	// }
}
