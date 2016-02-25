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

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": 123}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when no args", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c"}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when valid with extra properties", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"a", "b", "c", "d"}}}, ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"a"}}}, ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"a", 1, "a"}}}, ctx); errs == nil {
		t.Error("Should fail when invalid types supplied", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"map", "key", "subkey"}}}, ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	if _, errs := validateFindInMap(parse.FindInMap{map[string]interface{}{"Fn::FindInMap": []interface{}{"map", parse.Ref{map[string]interface{}{"Ref": "MyResource"}}, "subkey"}}}, ctx); errs != nil {
		t.Error("Should short circuit and pass when ref used", errs)
	}
}
