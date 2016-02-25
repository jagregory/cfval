package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestFindInMap(t *testing.T) {
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
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": 123}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when no args", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c"}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when valid with extra properties", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c", "d"}}}, ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a"}}}, ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", 1, "a"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid types supplied", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"map", "key", "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	ref := parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "MyResource"}}
	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{ref, "key", "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when Ref used as MapName", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"map", ref, "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when Ref used as TopLevelKey", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"map", "key", ref}}}, ctx); errs != nil {
		t.Error("Should pass when Ref used as SecondLevelKey", errs)
	}

	fim := parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c"}}}
	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{fim, "key", "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when Fn::FindInMap used as MapName", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"map", fim, "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when Fn::FindInMap used as TopLevelKey", errs)
	}

	if _, errs := validateFindInMap(parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": []interface{}{"map", "key", fim}}}, ctx); errs != nil {
		t.Error("Should pass when Fn::FindInMap used as SecondLevelKey", errs)
	}
}
