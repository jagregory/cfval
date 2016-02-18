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
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := []string{}

	schema := Schema{
		Type: ValueString,
	}

	if _, errs := schema.Validate("abc", self, nil, nil, ctx); errs != nil {
		t.Error("Should pass when value is correct type")
	}

	if _, errs := schema.Validate(123, self, nil, nil, ctx); errs == nil {
		t.Error("Should fail when value is incorrect type")
	}
}

func TestSchemaArrayValidation(t *testing.T) {
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := []string{}

	schema := Schema{
		Type:  ValueString,
		Array: true,
	}

	if _, errs := schema.Validate([]interface{}{"abc"}, self, nil, nil, ctx); errs != nil {
		t.Error("Should pass when value is an array of the correct type")
	}

	if _, errs := schema.Validate([]interface{}{"abc", 123}, self, nil, nil, ctx); errs == nil {
		t.Error("Should fail when value is a mixed array")
	}

	if _, errs := schema.Validate([]interface{}{123}, self, nil, nil, ctx); errs == nil {
		t.Error("Should fail when value is an incorrect array")
	}
}

func TestSchemaCustomValidation(t *testing.T) {
	res := Resource{
		ReturnValue: Schema{
			Type: ValueNumber,
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

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
	ctx := []string{}

	schema := Schema{
		Type:         ValueNumber,
		ValidateFunc: IntegerRangeValidate(10, 15),
	}

	if _, errs := schema.Validate(float64(12), self, template, definitions, ctx); errs != nil {
		t.Error("Should run custom validation when type is correct", errs)
	}

	if _, errs := schema.Validate(float64(20), self, template, definitions, ctx); errs == nil {
		t.Error("Should run custom validation when type is correct")
	}

	if _, errs := schema.Validate("abc", self, template, definitions, ctx); errs != nil && errs[0].Message != "Property has invalid type string, expected: ValueNumber" {
		t.Error("Should not run validation when type is correct", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "abc"}, self, template, definitions, ctx); errs != nil {
		t.Error("Should not run validation with Ref", errs)
	}
}
