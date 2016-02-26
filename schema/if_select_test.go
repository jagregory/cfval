package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSelect(t *testing.T) {
	template := &parse.Template{
		Conditions: map[string]parse.Condition{
			"Condition": parse.Condition{},
		},

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
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	if _, errs := validateSelect(IF(parse.FnSelect)(nil), ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{}), ctx); errs == nil {
		t.Error("Should fail when empty argument supplied", errs)
	}

	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when key missing", errs)
	}

	validArray := []interface{}{"a", "b", "c"}
	if _, errs := validateSelect(parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), validArray}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when extra keys", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)(123), ctx); errs == nil {
		t.Error("Should fail when wrong type used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), validArray, 2}), ctx); errs == nil {
		t.Error("Should fail when wrong number of args", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), validArray}), ctx); errs != nil {
		t.Error("Should pass when valid indexer and array used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(10), validArray}), ctx); errs == nil {
		t.Error("Should fail when out of range index used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(-1), validArray}), ctx); errs == nil {
		t.Error("Should fail when negative index used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), nil}), ctx); errs == nil {
		t.Error("Should fail when nil array used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), []interface{}{"one", nil, "three"}}), ctx); errs == nil {
		t.Error("Should fail when array has nils used", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{ExampleValidIFs[parse.FnRef](), validArray}), ctx); errs != nil {
		t.Error("Should pass with Ref in index", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{ExampleValidIFs[parse.FnFindInMap](), validArray}), ctx); errs != nil {
		t.Error("Should pass with FindInMap in index", errs)
	}

	invalidIndexFns := parse.AllIntrinsicFunctions.
		Except(parse.FnFindInMap, parse.FnRef)
	for _, fn := range invalidIndexFns {
		if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{ExampleValidIFs[fn](), validArray}), ctx); errs == nil {
			t.Errorf("Should fail with %s in index: %s", fn, errs)
		}
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[parse.FnRef]()}), ctx); errs != nil {
		t.Error("Should pass with Ref in array", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[parse.FnFindInMap]()}), ctx); errs != nil {
		t.Error("Should pass with FindInMap in array", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[parse.FnGetAtt]()}), ctx); errs != nil {
		t.Error("Should pass with GetAtt in array", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[parse.FnGetAZs]()}), ctx); errs != nil {
		t.Error("Should pass with GetAZs in array", errs)
	}

	if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[parse.FnIf]()}), ctx); errs != nil {
		t.Error("Should pass with If in array", errs)
	}

	invalidArrayFns := parse.AllIntrinsicFunctions.
		Except(parse.FnFindInMap, parse.FnGetAZs, parse.FnGetAtt, parse.FnIf, parse.FnRef)
	for _, fn := range invalidArrayFns {
		if _, errs := validateSelect(IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[fn]()}), ctx); errs == nil {
			t.Errorf("Should fail with %s as array: %s", fn, errs)
		}
	}
}
