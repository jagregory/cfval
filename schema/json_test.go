package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJSONValidation(t *testing.T) {
	p := Schema{
		Type: JSON,
	}

	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"Resource1": parse.TemplateResource{
				Type: "ResourceA",
			},
			"Resource2": parse.TemplateResource{
				Type: "ResourceB",
			},
		},
	}
	tr := parse.TemplateResource{}
	currentResource := ResourceWithDefinition{tr, Resource{}}

	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"ResourceA": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},

		"ResourceB": Resource{
			ReturnValue: Schema{
				Type: ValueNumber,
			},
		},
	}), currentResource, p, ValidationOptions{})

	validRefs := map[string]interface{}{
		"One": map[string]interface{}{
			"Value": map[string]interface{}{
				"Ref": "Resource1",
			},
		},
		"Two": []interface{}{
			map[string]interface{}{"Ref": "Resource2"},
		},
	}

	invalidRefs := map[string]interface{}{
		"One": map[string]interface{}{
			"Value": map[string]interface{}{
				"Ref": "Resource9",
			},
		},
		"Two": []interface{}{
			map[string]interface{}{"Ref": "Resource10"},
		},
	}

	if _, errs := JSON.Validate(validRefs, ctx); errs != nil {
		t.Error("Should pass with valid refs", errs)
	}

	if _, errs := JSON.Validate(invalidRefs, ctx); errs == nil {
		t.Error("Should fail with invalid refs")
	}
}
