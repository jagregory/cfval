package schema

import (
	"fmt"
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

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnRef)(123), false, "invalid type used for target"},
		IFScenario{IF(parse.FnRef)(nil), false, "nil used for target"},
		IFScenario{parse.IntrinsicFunction{"Ref", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": []interface{}{"delim", []interface{}{"a", "b"}}, "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnRef)("Resource1"), false, "valid resource ref with Unknown ref type"},
		IFScenario{IF(parse.FnRef)("NoTypeParameter"), false, "valid parameter ref with Unknown ref type"},
		IFScenario{IF(parse.FnRef)("Resource2"), true, "valid resource ref with matching types"},
		IFScenario{IF(parse.FnRef)("StringParameter"), true, "valid parameter ref with matching types"},
		IFScenario{IF(parse.FnRef)("AWS::StackName"), true, "valid pseudo-parameter ref with matching types"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnRef)(ExampleValidIFs[fn]()), false, fmt.Sprintf("%s as target", fn)})
	}

	for i, s := range scenarios {
		errs := validateRef(s.fn, stringContext)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}

	if errs := validateRef(IF(parse.FnRef)("Resource2"), numberContext); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if errs := validateRef(IF(parse.FnRef)("StringParameter"), numberContext); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if errs := validateRef(IF(parse.FnRef)("AWS::StackName"), numberContext); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if errs := validateRef(IF(parse.FnRef)("ListInstanceIdParameter"), listInstanceIDContext); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if errs := validateRef(IF(parse.FnRef)("AWS::NoValue"), stringContext); errs != nil {
		t.Error("Should pass on AWS::NoValue being assigned to anything", errs)
	}

	if errs := validateRef(IF(parse.FnRef)("AWS::NoValue"), numberContext); errs != nil {
		t.Error("Should pass on AWS::NoValue being assigned to anything", errs)
	}

	if errs := validateRef(IF(parse.FnRef)("AWS::NoValue"), listInstanceIDContext); errs != nil {
		t.Error("Should pass on AWS::NoValue being assigned to anything", errs)
	}
}
