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

	property := Schema{Type: ValueString}
	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"good": parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	self := ResourceWithDefinition{
		parse.TemplateResource{},
		res,
	}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}), self, property, ValidationOptions{})

	if _, errs := ValueString.Validate("abc", ctx); errs != nil {
		t.Error("Should pass with valid String")
	}

	if _, errs := ValueString.Validate(100, ctx); errs == nil {
		t.Error("Should fail with non-String")
	}

	if _, errs := ValueString.Validate(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "bad"}}, ctx); errs == nil {
		t.Error("Should fail with invalid ref")
	}

	result, errs := ValueString.Validate(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "good"}}, ctx)
	if errs != nil {
		t.Error("Should pass with valid ref", errs)
	}
	if result != reporting.ValidateAbort {
		t.Error("Should always abort validation when something is a intrinsic function but isn't valid - this prevents further validation on something which looks like a complex type")
	}

	// TODO: test other intrinsic functions are correctly handled by valuetype

	if _, errs := ValueString.Validate(map[string]interface{}{"something": "else"}, ctx); errs == nil {
		t.Error("Should fail with non-intrinsic function map")
	}
}
