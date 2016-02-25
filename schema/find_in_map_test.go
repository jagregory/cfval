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

	if _, errs := NewFindInMap(nil).Validate(ctx); errs == nil {
		t.Error("Should fail when no arguments supplied", errs)
	}

	if _, errs := NewFindInMap([]interface{}{"a", "b", "c", "d"}).Validate(ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := NewFindInMap([]interface{}{"a"}).Validate(ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := NewFindInMap([]interface{}{"a", 1, "a"}).Validate(ctx); errs == nil {
		t.Error("Should fail when invalid types supplied", errs)
	}

	if _, errs := NewFindInMap([]interface{}{"map", "key", "subkey"}).Validate(ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	if _, errs := NewFindInMap([]interface{}{"map", parse.Ref{map[string]interface{}{"Ref": "MyResource"}}, "subkey"}).Validate(ctx); errs != nil {
		t.Error("Should short circuit and pass when ref used", errs)
	}
}
