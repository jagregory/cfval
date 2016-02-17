package schema

import (
	"testing"

	"github.com/jagregory/cfval/constraints"
)

func TestNestedResourceConstraints(t *testing.T) {
	data := func(properties map[string]interface{}) TemplateResource {
		return TemplateResource{
			Definition: Resource{
				Properties: Properties{
					"Nested": Schema{
						Type: NestedResource{
							Properties: Properties{
								"One": Schema{
									Type: ValueString,
								},

								"Two": Schema{
									Type:     ValueString,
									Required: constraints.PropertyExists("One"),
								},
							},
						},
					},
				},
			},

			Properties: properties,
		}
	}

	twoMissing := map[string]interface{}{
		"Nested": map[string]interface{}{
			"One": "abc",
		},
	}
	if _, errs := data(twoMissing).Validate([]string{}); errs == nil {
		t.Error("Should fail with missing Two parameter")
	}

	oneInWrongPace := map[string]interface{}{
		"One":    "abc",
		"Nested": map[string]interface{}{},
	}
	if _, errs := data(oneInWrongPace).Validate([]string{}); errs == nil {
		t.Error("Should fail with missing Two parameter")
	}

	allFine := map[string]interface{}{
		"Nested": map[string]interface{}{
			"One": "abc",
			"Two": "abc",
		},
	}
	if _, errs := data(allFine).Validate([]string{}); errs != nil {
		t.Error("Should pass with One and Two", errs)
	}
}
