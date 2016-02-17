package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJSONValidation(t *testing.T) {
	p := Schema{
		Type: JSON,
	}

	definitions := NewResourceDefinitions(map[string]func() Resource{
		"ResourceA": func() Resource {
			return Resource{
				ReturnValue: Schema{
					Type: ValueString,
				},
			}
		},

		"ResourceB": func() Resource {
			return Resource{
				ReturnValue: Schema{
					Type: ValueNumber,
				},
			}
		},
	})

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
	tr := parse.TemplateResource{
		Tmpl: template,
	}
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

	if _, errs := JSON.Validate(p, validRefs, tr, definitions, ctx); errs != nil {
		t.Error("Should pass with valid refs", errs)
	}

	if _, errs := JSON.Validate(p, invalidRefs, tr, definitions, ctx); errs == nil {
		t.Error("Should fail with invalid refs")
	}
}
