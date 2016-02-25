package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSelect(t *testing.T) {
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

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": nil}}, ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{}}}, ctx); errs == nil {
		t.Error("Should fail when empty argument supplied", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when key missing", errs)
	}

	validArray := []interface{}{"a", "b", "c"}
	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), validArray}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when extra keys", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": 123}}, ctx); errs == nil {
		t.Error("Should fail when wrong type used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), validArray, 2}}}, ctx); errs == nil {
		t.Error("Should fail when wrong number of args", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), validArray}}}, ctx); errs != nil {
		t.Error("Should pass when valid indexer and array used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(10), validArray}}}, ctx); errs == nil {
		t.Error("Should fail when out of range index used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(-1), validArray}}}, ctx); errs == nil {
		t.Error("Should fail when negative index used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), nil}}}, ctx); errs == nil {
		t.Error("Should fail when nil array used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), []interface{}{"one", nil, "three"}}}}, ctx); errs == nil {
		t.Error("Should fail when array has nils used", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "MyResource"}}, validArray}}}, ctx); errs != nil {
		t.Error("Should pass with Ref in index", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c"}}}, validArray}}}, ctx); errs != nil {
		t.Error("Should pass with FindInMap in index", errs)
	}

	invalidIndexFns := parse.AllIntrinsicFunctions.
		Except(parse.FnFindInMap, parse.FnRef)
	for _, fn := range invalidIndexFns {
		if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}, validArray}}}, ctx); errs == nil {
			t.Errorf("Should fail with %s in index: %s", fn, errs)
		}
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "MyResource"}}}}}, ctx); errs != nil {
		t.Error("Should pass with Ref in array", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c"}}}}}}, ctx); errs != nil {
		t.Error("Should pass with FindInMap in array", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "InstanceId"}}}}}}, ctx); errs != nil {
		t.Error("Should pass with GetAtt in array", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": ""}}}}}, ctx); errs != nil {
		t.Error("Should pass with GetAZs in array", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{"Fn::If", map[string]interface{}{"Fn::If": ""}}}}}, ctx); errs != nil {
		t.Error("Should pass with If in array", errs)
	}

	invalidArrayFns := parse.AllIntrinsicFunctions.
		Except(parse.FnFindInMap, parse.FnGetAZs, parse.FnGetAtt, parse.FnIf, parse.FnRef)
	for _, fn := range invalidArrayFns {
		if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}}}}, ctx); errs == nil {
			t.Errorf("Should fail with %s as array: %s", fn, errs)
		}
	}
}
