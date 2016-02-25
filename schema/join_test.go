package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJoin(t *testing.T) {
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

				"Name": Schema{
					Type: ValueString,
				},
			},

			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": 123}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when no args", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when valid with extra properties", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"a", "b", "c", "d"}}}, ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"a"}}}, ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{1, []interface{}{"a", "b"}}}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for delimeter", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"a", 1}}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"a", []interface{}{1, "b"}}}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used in args", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"d", []interface{}{"a", "b"}}}}, ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	if _, errs := validateJoin(parse.Join{map[string]interface{}{"Fn::Join": []interface{}{"d", []interface{}{parse.Ref{map[string]interface{}{"Ref": "MyResource"}}, "b", "c"}}}}, ctx); errs != nil {
		t.Error("Should short circuit and pass when ref used", errs)
	}
}
