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
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})
	numberContext := NewPropertyContext(stringContext, Schema{Type: ValueNumber})
	listInstanceIDContext := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: Multiple(InstanceID)}, ValidationOptions{})

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": 123}}, stringContext); errs == nil {
		t.Error("Should fail on ref with invalid target type")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{}}, stringContext); errs == nil {
		t.Error("Should fail on ref with no target")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "Resource2", "blah": "blah"}}, stringContext); errs == nil {
		t.Error("Should fail on valid ref with extra properties")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "Resource1"}}, stringContext); errs == nil {
		t.Error("Should fail on valid resource ref with Unknown ref type")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "Resource2"}}, stringContext); errs != nil {
		t.Error("Should pass on valid resource ref with matching types", errs)
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "Resource2"}}, numberContext); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "NoTypeParameter"}}, stringContext); errs == nil {
		t.Error("Should fail on valid parameter ref with Unknown ref type", errs)
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "StringParameter"}}, stringContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "StringParameter"}}, numberContext); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "AWS::StackName"}}, stringContext); errs != nil {
		t.Error("Should pass on valid pseudo-parameter ref with matching types", errs)
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "AWS::StackName"}}, numberContext); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "StringParameter"}}, stringContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": "ListInstanceIdParameter"}}, listInstanceIDContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	invalidFns := parse.AllIntrinsicFunctions
	for _, fn := range invalidFns {
		if _, errs := validateRef(parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": parse.IntrinsicFunction{fn, map[string]interface{}{string(fn): "MyResource"}}}}, stringContext); errs == nil {
			t.Errorf("Should fail with %s as target: %s", fn, errs)
		}
	}
}
