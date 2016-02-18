package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSchemaTargetType(t *testing.T) {
	if (Schema{Type: ValueNumber}).TargetType() != ValueNumber {
		t.Error("Schema TargetType should match Type")
	}

	if (Schema{}).TargetType() != nil {
		t.Error("Schema without Type should return nil for TargetType")
	}
}

func TestSchemaTypeValidation(t *testing.T) {
	template := &parse.Template{}
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(nil), self, Schema{})
	schema := Schema{Type: ValueString}

	if _, errs := schema.Validate("abc", ctx); errs != nil {
		t.Error("Should pass when value is correct type")
	}

	if _, errs := schema.Validate(123, ctx); errs == nil {
		t.Error("Should fail when value is incorrect type")
	}
}

func TestSchemaArrayValidation(t *testing.T) {
	template := &parse.Template{}
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(nil), self, Schema{})

	schema := Schema{
		Type:  ValueString,
		Array: true,
	}

	if _, errs := schema.Validate([]interface{}{"abc"}, ctx); errs != nil {
		t.Error("Should pass when value is an array of the correct type")
	}

	if _, errs := schema.Validate([]interface{}{"abc", 123}, ctx); errs == nil {
		t.Error("Should fail when value is a mixed array")
	}

	if _, errs := schema.Validate([]interface{}{123}, ctx); errs == nil {
		t.Error("Should fail when value is an incorrect array")
	}
}

func TestSchemaCustomValidation(t *testing.T) {
	res := Resource{
		ReturnValue: Schema{
			Type: ValueNumber,
		},
	}

	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"abc": &parse.TemplateResource{
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
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}), self, Schema{})

	schema := Schema{
		Type:         ValueNumber,
		ValidateFunc: IntegerRangeValidate(10, 15),
	}

	if _, errs := schema.Validate(float64(12), ctx); errs != nil {
		t.Error("Should run custom validation when type is correct", errs)
	}

	if _, errs := schema.Validate(float64(20), ctx); errs == nil {
		t.Error("Should run custom validation when type is correct")
	}

	if _, errs := schema.Validate("abc", ctx); errs != nil && errs[0].Message != "Property has invalid type string, expected: ValueNumber" {
		t.Error("Should not run validation when type is correct", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "abc"}, ctx); errs != nil {
		t.Error("Should not run validation with Ref", errs)
	}
}
