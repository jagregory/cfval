package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestRefValidate(t *testing.T) {
	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"Resource1": parse.TemplateResource{},
			"Resource2": parse.TemplateResource{
				Type: "TestResource",
			},
		},
		Parameters: map[string]parse.Parameter{
			"NoTypeParameter": parse.Parameter{},
			"StringParameter": parse.Parameter{
				Type: "String",
			},
			"ListInstanceIdParameter": parse.Parameter{
				Type: "List<AWS::EC2::Instance::Id>",
			},
		},
	}

	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	stringContext := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString})
	numberContext := NewPropertyContext(stringContext, Schema{Type: ValueNumber})
	listInstanceIDContext := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: Multiple(InstanceID)})

	if _, errs := NewRef("Resource1").Validate(stringContext); errs == nil {
		t.Error("Should fail on valid resource ref with Unknown ref type")
	}

	if _, errs := NewRef("Resource2").Validate(stringContext); errs != nil {
		t.Error("Should pass on valid resource ref with matching types", errs)
	}

	if _, errs := NewRef("Resource2").Validate(numberContext); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if _, errs := NewRef("NoTypeParameter").Validate(stringContext); errs == nil {
		t.Error("Should fail on valid parameter ref with Unknown ref type", errs)
	}

	if _, errs := NewRef("StringParameter").Validate(stringContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := NewRef("StringParameter").Validate(numberContext); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if _, errs := NewRef("AWS::StackName").Validate(stringContext); errs != nil {
		t.Error("Should pass on valid pseudo-parameter ref with matching types", errs)
	}

	if _, errs := NewRef("AWS::StackName").Validate(numberContext); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if _, errs := NewRef("StringParameter").Validate(stringContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := NewRef("ListInstanceIdParameter").Validate(listInstanceIDContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := NewRef("invalid").Validate(stringContext); errs == nil {
		t.Error("Should fail on invalid ref")
	}
}
