package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJSONValidation(t *testing.T) {
	p := Schema{
		Type: JSON,
	}

	definitions := NewResourceDefinitions(map[string]Resource{
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
	})

	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"Resource1": &parse.TemplateResource{
				Type: "ResourceA",
			},
			"Resource2": &parse.TemplateResource{
				Type: "ResourceB",
			},
		},
	}
	tr := parse.TemplateResource{
		Tmpl: template,
	}
	currentResource := ResourceWithDefinition{tr, Resource{}}

	ctx := []string{}

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

	if _, errs := JSON.Validate(p, validRefs, currentResource, template, definitions, ctx); errs != nil {
		t.Error("Should pass with valid refs", errs)
	}

	if _, errs := JSON.Validate(p, invalidRefs, currentResource, template, definitions, ctx); errs == nil {
		t.Error("Should fail with invalid refs")
	}
}
