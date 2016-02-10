package schema

import (
	"testing"

	"github.com/jagregory/cfval/reporting"
)

func TestValueTypeValidation(t *testing.T) {
	property := Schema{Type: ValueString}
	self := TemplateResource{
		template: &Template{
			Resources: map[string]TemplateResource{
				"good": TemplateResource{
					Definition: Resource{
						ReturnValue: Schema{
							Type: ValueString,
						},
					},
				},
			},
		},
	}
	ctx := []string{}

	if _, errs := ValueString.Validate(property, "abc", self, ctx); errs != nil {
		t.Error("Should pass with valid String")
	}

	if _, errs := ValueString.Validate(property, 100, self, ctx); errs == nil {
		t.Error("Should fail with non-String")
	}

	if _, errs := ValueString.Validate(property, map[string]interface{}{"Ref": "bad"}, self, ctx); errs == nil {
		t.Error("Should fail with invalid ref")
	}

	result, errs := ValueString.Validate(property, map[string]interface{}{"Ref": "good"}, self, ctx)
	if errs != nil {
		t.Error("Should pass with valid ref", errs)
	}
	if result != reporting.ValidateAbort {
		t.Error("Should always abort validation when something is a builtin but isn't valid - this prevents further validation on something which looks like a complex type")
	}

	// TODO: test other builtins are correctly handled by valuetype

	if _, errs := ValueString.Validate(property, map[string]interface{}{"something": "else"}, self, ctx); errs == nil {
		t.Error("Should fail with non-builtin map")
	}
}
