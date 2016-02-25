package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestGetAZs(t *testing.T) {
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

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": nil}}, ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": ""}}, ctx); errs != nil {
		t.Error("Should pass when empty argument supplied", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when key missing", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": "ap-southeast-2", "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when extra keys", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": 123}}, ctx); errs == nil {
		t.Error("Should fail when wrong type used", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": "ap-southeast-2"}}, ctx); errs != nil {
		t.Error("Should pass when region used", errs)
	}

	if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "MyResource"}}}}, ctx); errs != nil {
		t.Error("Should pass with nested Ref", errs)
	}

	invalidFns := []parse.IntrinsicFunctionSignature{
		parse.FnAnd,
		parse.FnBase64,
		parse.FnEquals,
		parse.FnFindInMap,
		parse.FnGetAtt,
		parse.FnGetAZs,
		parse.FnIf,
		parse.FnJoin,
		parse.FnNot,
		parse.FnOr,
		parse.FnSelect,
	}
	for _, fn := range invalidFns {
		if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}}}, ctx); errs == nil {
			t.Errorf("Should fail with nested %s: %s", fn, errs)
		}
	}

	// TODO: region validation
	// if _, errs := validateGetAZs(parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": "ap-southeast-9"}}, ctx); errs == nil {
	// 	t.Error("Should fail when invalid region used", errs)
	// }
}
