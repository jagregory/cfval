package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestGetAtt(t *testing.T) {
	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			Attributes: Properties{
				"InstanceId": Schema{
					Type: InstanceID,
				},

				"ListInstanceId": Schema{
					Type:  InstanceID,
					Array: true,
				},

				"Name": Schema{
					Type: ValueString,
				},
			},
		},
	}), currentResource, Schema{Type: InstanceID})
	listCtx := NewPropertyContext(ctx, Schema{Type: InstanceID, Array: true})

	if _, errs := NewGetAtt(nil).Validate(ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"a", "b", "c"}).Validate(ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"a"}).Validate(ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"UnknownResource", "prop"}).Validate(ctx); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"MyResource", "BadProp"}).Validate(ctx); errs == nil {
		t.Error("Should fail when invalid property used for type of resource", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"MyResource", "Name"}).Validate(ctx); errs == nil {
		t.Error("Should fail when valid property of wrong type", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"MyResource", "InstanceId"}).Validate(ctx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}

	if _, errs := NewGetAtt([]interface{}{"MyResource", "ListInstanceId"}).Validate(listCtx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}
}
