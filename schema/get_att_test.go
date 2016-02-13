package schema

import "testing"

func TestGetAtt(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{
				Definition: Resource{
					Attributes: Properties{
						"InstanceId": Schema{
							Type: InstanceID,
						},

						"Name": Schema{
							Type: ValueString,
						},
					},
				},
			},
		},
	}
	context := []string{}
	prop := Schema{
		Type: InstanceID,
	}

	if _, errs := NewGetAtt(prop, nil).Validate(template, context); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"a", "b", "c"}).Validate(template, context); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"a"}).Validate(template, context); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"UnknownResource", "prop"}).Validate(template, context); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "BadProp"}).Validate(template, context); errs == nil {
		t.Error("Should fail when invalid property used for type of resource", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "Name"}).Validate(template, context); errs == nil {
		t.Error("Should fail when valid property of wrong type", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "InstanceId"}).Validate(template, context); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}
}
