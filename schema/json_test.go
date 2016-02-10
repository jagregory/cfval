package schema

import "testing"

func TestJSONValidation(t *testing.T) {
	p := Schema{
		Type: JSON,
	}
	template := &Template{
		Resources: map[string]TemplateResource{
			"Resource1": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: ValueString,
					},
				},
			},
			"Resource2": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: ValueNumber,
					},
				},
			},
		},
	}
	tr := TemplateResource{
		template: template,
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

	if _, errs := JSON.Validate(p, validRefs, tr, ctx); errs != nil {
		t.Error("Should pass with valid refs", errs)
	}

	if _, errs := JSON.Validate(p, invalidRefs, tr, ctx); errs == nil {
		t.Error("Should fail with invalid refs")
	}
}
