package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func TestValueTypeValidation(t *testing.T) {
	res := Resource{
		ReturnValue: Schema{
			Type: ValueString,
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	property := Schema{Type: ValueString}
	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"good": &parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	self := ResourceWithDefinition{
		parse.TemplateResource{
			Tmpl: template,
		},
		res,
	}
	ctx := []string{}

	if _, errs := ValueString.Validate(property, "abc", self, template, definitions, ctx); errs != nil {
		t.Error("Should pass with valid String")
	}

	if _, errs := ValueString.Validate(property, 100, self, template, definitions, ctx); errs == nil {
		t.Error("Should fail with non-String")
	}

	if _, errs := ValueString.Validate(property, map[string]interface{}{"Ref": "bad"}, self, template, definitions, ctx); errs == nil {
		t.Error("Should fail with invalid ref")
	}

	result, errs := ValueString.Validate(property, map[string]interface{}{"Ref": "good"}, self, template, definitions, ctx)
	if errs != nil {
		t.Error("Should pass with valid ref", errs)
	}
	if result != reporting.ValidateAbort {
		t.Error("Should always abort validation when something is a builtin but isn't valid - this prevents further validation on something which looks like a complex type")
	}

	// TODO: test other builtins are correctly handled by valuetype

	if _, errs := ValueString.Validate(property, map[string]interface{}{"something": "else"}, self, template, definitions, ctx); errs == nil {
		t.Error("Should fail with non-builtin map")
	}
}
