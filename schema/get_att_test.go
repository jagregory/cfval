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
					Type: Multiple(InstanceID),
				},

				"Name": Schema{
					Type: ValueString,
				},
			},
		},
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})
	listCtx := NewPropertyContext(ctx, Schema{Type: Multiple(InstanceID)})

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{}}}, ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when key missing", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"a", "b"}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when extra keys", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"a", "b", "c"}}}, ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"a"}}}, ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"UnknownResource", "prop"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "BadProp"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid property used for type of resource", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "Name"}}}, ctx); errs == nil {
		t.Error("Should fail when valid property of wrong type", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "InstanceId"}}}, ctx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}

	if _, errs := validateGetAtt(parse.GetAtt{map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "ListInstanceId"}}}, listCtx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}
}
