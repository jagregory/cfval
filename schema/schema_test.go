package schema

import "testing"

func TestGetAtt(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{},
		},
	}
	tr := TemplateResource{
		Template: template,
	}
	context := []string{}

	if ok, errs := validateGetAtt(nil, tr, context); ok {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if ok, errs := validateGetAtt([]interface{}{"a", "b", "c"}, tr, context); ok {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if ok, errs := validateGetAtt([]interface{}{"a"}, tr, context); ok {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if ok, errs := validateGetAtt([]interface{}{"UnknownResource", "prop"}, tr, context); ok {
		t.Error("Should fail when invalid resource used", errs)
	}

	// TODO: this test will eventually fail when we implement GetAtt prop validation
	// uncomment the following tests later
	if ok, errs := validateGetAtt([]interface{}{"MyResource", "prop"}, tr, context); !ok {
		t.Error("Should pass when valid resource used", errs)
	}

	// if ok, errs := validateGetAtt([]interface{}{"MyResource", "BadProp"}, tr, context); ok {
	// 	t.Error("Should fail when invalid property used for type of resource", errs)
	// }
	//
	// if ok, errs := validateGetAtt([]interface{}{"MyResource", "InstanceId"}, tr, context); !ok {
	// 	t.Error("Should pass when valid property used for type of resource", errs)
	// }
}
