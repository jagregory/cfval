package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestGetAtt(t *testing.T) {
	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			Attributes: Properties{
				"InstanceId": Schema{
					Type: InstanceID,
				},

				"Name": Schema{
					Type: ValueString,
				},
			},
		},
	})

	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"MyResource": &parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	path := []string{}
	prop := Schema{
		Type: InstanceID,
	}

	if _, errs := NewGetAtt(prop, nil).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"a", "b", "c"}).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"a"}).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"UnknownResource", "prop"}).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "BadProp"}).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when invalid property used for type of resource", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "Name"}).Validate(template, definitions, path); errs == nil {
		t.Error("Should fail when valid property of wrong type", errs)
	}

	if _, errs := NewGetAtt(prop, []interface{}{"MyResource", "InstanceId"}).Validate(template, definitions, path); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}
}
