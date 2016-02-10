package schema

import "testing"

func TestSchemaTargetType(t *testing.T) {
	if (Schema{Type: ValueNumber}).TargetType() != ValueNumber {
		t.Error("Schema TargetType should match Type")
	}

	if (Schema{}).TargetType() != ValueUnknown {
		t.Error("Schema without Type should return TypeUnknown for TargetType")
	}
}

func TestSchemaRequiredValidation(t *testing.T) {
	self := TemplateResource{}
	ctx := []string{}

	notRequired := Schema{
		Type: ValueString,
	}
	required := Schema{
		Type:     ValueString,
		Required: true,
	}

	if _, errs := notRequired.Validate(nil, self, ctx); errs != nil {
		t.Error("Should pass when a not required field is nil")
	}

	if _, errs := notRequired.Validate("abc", self, ctx); errs != nil {
		t.Error("Should pass when a not required field has a value")
	}

	if _, errs := required.Validate(nil, self, ctx); errs == nil {
		t.Error("Should fail when a required field is nil")
	}

	if _, errs := required.Validate("abc", self, ctx); errs != nil {
		t.Error("Should pass when a required field has a value")
	}
}

func TestSchemaTypeValidation(t *testing.T) {
	self := TemplateResource{}
	ctx := []string{}

	schema := Schema{
		Type: ValueString,
	}

	if _, errs := schema.Validate("abc", self, ctx); errs != nil {
		t.Error("Should pass when value is correct type")
	}

	if _, errs := schema.Validate(123, self, ctx); errs == nil {
		t.Error("Should fail when value is incorrect type")
	}
}

func TestSchemaArrayValidation(t *testing.T) {
	self := TemplateResource{}
	ctx := []string{}

	schema := Schema{
		Type:  ValueString,
		Array: true,
	}

	if _, errs := schema.Validate([]interface{}{"abc"}, self, ctx); errs != nil {
		t.Error("Should pass when value is an array of the correct type")
	}

	if _, errs := schema.Validate([]interface{}{"abc", 123}, self, ctx); errs == nil {
		t.Error("Should fail when value is a mixed array")
	}

	if _, errs := schema.Validate([]interface{}{123}, self, ctx); errs == nil {
		t.Error("Should fail when value is an incorrect array")
	}
}

func TestSchemaCustomValidation(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"abc": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: ValueNumber,
					},
				},
			},
		},
	}
	self := TemplateResource{
		template: template,
	}
	ctx := []string{}

	schema := Schema{
		Type:         ValueNumber,
		ValidateFunc: IntegerRangeValidate(10, 15),
	}

	if _, errs := schema.Validate(float64(12), self, ctx); errs != nil {
		t.Error("Should run custom validation when type is correct", errs)
	}

	if _, errs := schema.Validate(float64(20), self, ctx); errs == nil {
		t.Error("Should run custom validation when type is correct")
	}

	if _, errs := schema.Validate("abc", self, ctx); errs != nil && errs[0].Message != "Property has invalid type string, expected: ValueNumber" {
		t.Error("Should not run validation when type is correct", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "abc"}, self, ctx); errs != nil {
		t.Error("Should not run validation with Ref", errs)
	}
}
