package schema

import (
	"testing"

	"github.com/jagregory/cfval/reporting"
)

type coercion struct{ from, to PropertyType }

var validCoercions = []coercion{
	coercion{from: KeyName, to: ValueString},
	// coercion{from: ValueString, to: KeyName},
	// coercion{from: ValueNumber, to: ValueString},
	// coercion{from: ValueString, to: ValueNumber},
}

func TestValidCoercions(t *testing.T) {
	for _, c := range validCoercions {
		template := &Template{}
		template.Resources = map[string]TemplateResource{
			"Source": TemplateResource{
				template: template,
				Definition: Resource{
					ReturnValue: Schema{
						Type: c.from,
					},
				},
			},

			"TestResource": TemplateResource{
				template: template,
				Definition: Resource{
					Properties: map[string]Schema{
						"Destination": Schema{
							Type: c.to,
						},
					},
				},
				Properties: map[string]interface{}{
					"Destination": map[string]interface{}{"Ref": "Source"},
				},
			},
		}

		_, errs := template.Validate()

		if errs != nil {
			t.Errorf("Should coerce %s into %s.\nFailures:\n%s", c.from, c.to, errs)
		}
	}
}

func testHasFailure(failures reporting.Failures, message string) (int, bool) {
	for i, failure := range failures {
		if failure.Message == message {
			return i, true
		}
	}

	return 0, false
}
