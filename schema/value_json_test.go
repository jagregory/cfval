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

	validRef := map[string]interface{}{
		"One": map[string]interface{}{
			"Value": IF(parse.FnRef)("Resource1"),
		},
	}

	invalidRefs := map[string]interface{}{
		"One": map[string]interface{}{
			"Value": IF(parse.FnRef)("Resource9"),
		},
		"Two": []interface{}{IF(parse.FnRef)("Resource10")},
	}

	if _, errs := JSON.Validate(validRef, ctx); errs != nil {
		t.Errorf("Should pass with valid refs (errs: %s)", errs)
	}

	if _, errs := JSON.Validate(invalidRefs, ctx); errs == nil {
		t.Error("Should fail with invalid refs")
	}
}
