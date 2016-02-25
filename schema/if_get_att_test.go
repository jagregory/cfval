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

			ReturnValue: Schema{
				Type: InstanceID,
			},
		},
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})
	listCtx := NewPropertyContext(ctx, Schema{Type: Multiple(InstanceID)})

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{}}}, ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when key missing", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"a", "b"}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when extra keys", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"a", "b", "c"}}}, ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"a"}}}, ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"UnknownResource", "prop"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid resource used", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "BadProp"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid property used for type of resource", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "Name"}}}, ctx); errs == nil {
		t.Error("Should fail when valid property of wrong type", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "InstanceId"}}}, ctx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "ListInstanceId"}}}, listCtx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}

	invalidResourceFns := parse.AllIntrinsicFunctions
	for _, fn := range invalidResourceFns {
		if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}, "InstanceId"}}}, ctx); errs == nil {
			t.Errorf("Should fail with %s in Resource: %s", fn, errs)
		}
	}

	if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "MyResource"}}}}}, ctx); errs != nil {
		t.Errorf("Should pass with Ref in Attribute: %s", errs)
	}

	invalidAttributeFns := parse.AllIntrinsicFunctions.
		Except(parse.FnRef)
	for _, fn := range invalidAttributeFns {
		if _, errs := validateGetAtt(parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}}}}, ctx); errs == nil {
			t.Errorf("Should fail with %s in Attribute: %s", fn, errs)
		}
	}
}
