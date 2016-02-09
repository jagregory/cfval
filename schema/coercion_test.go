package schema

import (
	"testing"

	"github.com/jagregory/cfval/reporting"
)

func TestCoercionOfKeyNameToString(t *testing.T) {
	template := &Template{}
	template.Resources = map[string]TemplateResource{
		"KeyNameResource": TemplateResource{
			template: template,
			Definition: Resource{
				ReturnValue: Schema{
					Type: KeyName,
				},
			},
		},

		"StringResource": TemplateResource{
			template: template,
			Definition: Resource{
				ReturnValue: Schema{
					Type: ValueString,
				},
			},
		},

		"TestResource": TemplateResource{
			template: template,
			Definition: Resource{
				Properties: map[string]Schema{
					"StringProp": Schema{
						Type: ValueString,
					},

					"KeyNameProp": Schema{
						Type: KeyName,
					},
				},
			},
			Properties: map[string]interface{}{
				"StringProp":  map[string]interface{}{"Ref": "KeyNameResource"},
				"KeyNameProp": map[string]interface{}{"Ref": "StringResource"},
			},
		},
	}

	_, errs := template.Validate()

	if i, found := testHasFailure(errs, "Ref value of 'KeyNameResource' is KeyName but is being assigned to a ValueString property"); found {
		errs[i] = nil
		t.Error("Should coerce KeyName into String")
	}

	if i, found := testHasFailure(errs, "Ref value of 'StringResource' is ValueString but is being assigned to a KeyName property"); found {
		errs[i] = nil
		t.Error("Should coerce String into KeyName")
	}

	for _, err := range errs {
		if err != nil {
			t.Error("Unexpected error", err.Message)
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
