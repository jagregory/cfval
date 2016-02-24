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
	template := &parse.Template{
		Parameters: map[string]parse.Parameter{
			"StringParam": parse.Parameter{Type: "String"},
		},

		Resources: map[string]parse.TemplateResource{
			"Resource1": parse.TemplateResource{
				Type: "ResourceDef1",
			},
		},
	}
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"ResourceDef1": Resource{
			Attributes: map[string]Schema{
				"Name": Schema{
					Type: ValueString,
				},

				"ID": Schema{
					Type: JSON,
				},
			},
		},
	}), self, Schema{})
	schema := Schema{Type: ValueString}

	if _, errs := schema.Validate("abc", ctx); errs != nil {
		t.Error("Should pass when value is correct type", errs)
	}

	if _, errs := schema.Validate(123, ctx); errs == nil {
		t.Error("Should fail when value is incorrect type")
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "StringParam"}, ctx); errs != nil {
		t.Error("Should pass when Ref is correct type", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "NumberParam"}, ctx); errs == nil {
		t.Error("Should fail when Ref is incorrect type")
	}

	if _, errs := schema.Validate(map[string]interface{}{"Fn::GetAtt": []interface{}{"Resource1", "Name"}}, ctx); errs != nil {
		t.Error("Should pass when GetAtt is correct type", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Fn::GetAtt": []interface{}{"Resource1", "ID"}}, ctx); errs == nil {
		t.Error("Should fail when GetAtt is incorrect type")
	}
}

func TestSchemaArrayValidation(t *testing.T) {
	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"Target": parse.TemplateResource{
				Type: "ResourceDef1",
			},

			"ArrayTarget": parse.TemplateResource{
				Type: "ResourceDef2",
			},
		},
	}
	self := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"ResourceDef1": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},

		"ResourceDef2": Resource{
			ReturnValue: Schema{
				Type: Multiple(ValueString),
			},
		},
	}), self, Schema{})

	schema := Schema{
		Type: Multiple(ValueString),
	}

	if _, errs := schema.Validate([]interface{}{"abc"}, ctx); errs != nil {
		t.Error("Should pass when value is an array of the correct type", errs)
	}

	if _, errs := schema.Validate([]interface{}{"abc", 123}, ctx); errs == nil {
		t.Error("Should fail when value is a mixed array")
	}

	if _, errs := schema.Validate([]interface{}{123}, ctx); errs == nil {
		t.Error("Should fail when value is an incorrect array")
	}

	if _, errs := schema.Validate([]interface{}{"abc", map[string]interface{}{"Ref": "Target"}}, ctx); errs != nil {
		t.Error("Should pass when value is an array with Refs of the correct type", errs)
	}

	if _, errs := schema.Validate([]interface{}{"abc", map[string]interface{}{"Ref": "ArrayTarget"}}, ctx); errs == nil {
		t.Error("Should fail when value is an array with Refs of the wrong type")
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "ArrayTarget"}, ctx); errs != nil {
		t.Error("Should pass when value is a Ref of the correct type is used for whole value", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "Target"}, ctx); errs == nil {
		t.Error("Should fail when value is a Ref of the correct type is used for whole value")
	}
}

func TestSchemaCustomValidation(t *testing.T) {
	res := Resource{
		ReturnValue: Schema{
			Type: ValueNumber,
		},
	}

	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"abc": parse.TemplateResource{
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

	if _, errs := schema.Validate("abc", ctx); errs != nil && errs[0].Message != "Number used in String property" {
		t.Error("Should not run validation when type is correct", errs)
	}

	if _, errs := schema.Validate(map[string]interface{}{"Ref": "abc"}, ctx); errs != nil {
		t.Error("Should not run validation with Ref", errs)
	}
}
