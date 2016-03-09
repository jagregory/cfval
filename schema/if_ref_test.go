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
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnRef)(123), ValueString, false, "invalid type used for target"},
		IFScenario{IF(parse.FnRef)(nil), ValueString, false, "nil used for target"},
		IFScenario{parse.IntrinsicFunction{"Ref", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Ref", map[string]interface{}{"Ref": []interface{}{"delim", []interface{}{"a", "b"}}, "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnRef)("Resource1"), ValueString, false, "valid resource ref with Unknown ref type"},
		IFScenario{IF(parse.FnRef)("NoTypeParameter"), ValueString, false, "valid parameter ref with Unknown ref type"},
		IFScenario{IF(parse.FnRef)("Resource2"), ValueString, true, "valid resource ref with matching types"},
		IFScenario{IF(parse.FnRef)("StringParameter"), ValueString, true, "valid parameter ref with matching types"},
		IFScenario{IF(parse.FnRef)("AWS::StackName"), ValueString, true, "valid pseudo-parameter ref with matching types"},
		IFScenario{IF(parse.FnRef)("Resource2"), ValueNumber, false, "valid resource ref with non-matching types"},
		IFScenario{IF(parse.FnRef)("StringParameter"), ValueNumber, false, "valid parameter ref with non-matching types"},
		IFScenario{IF(parse.FnRef)("AWS::StackName"), ValueNumber, false, "valid pseudo-parameter ref with non-matching types"},
		IFScenario{IF(parse.FnRef)("ListInstanceIdParameter"), Multiple(InstanceID), true, "valid parameter ref with matching types"},
		IFScenario{IF(parse.FnRef)("AWS::NoValue"), ValueString, true, "AWS::NoValue being assigned to anything"},
		IFScenario{IF(parse.FnRef)("AWS::NoValue"), ValueNumber, true, "AWS::NoValue being assigned to anything"},
		IFScenario{IF(parse.FnRef)("AWS::NoValue"), Multiple(InstanceID), true, "AWS::NoValue being assigned to anything"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnRef)(ExampleValidIFs[fn]()), ValueString, false, fmt.Sprintf("%s as target", fn)})
	}

	scenarios.evaluate(t, validateRef, ctx)
}
