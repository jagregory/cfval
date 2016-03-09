package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestNot(t *testing.T) {
	template := &parse.Template{
		Conditions: map[string]parse.Condition{
			"Condition": parse.Condition{},
		},

		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			Attributes: Properties{
				"InstanceId": Schema{
					Type: InstanceID,
				},

				"Name": Schema{
					Type: ValueString,
				},
			},

			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnNot)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnNot)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Not", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Not", map[string]interface{}{"Fn::Not": "blah", "blah": "blah"}}, ValueString, false, "extra properties"},
	}

	validFns := []parse.IntrinsicFunctionSignature{
		parse.FnAnd,
		parse.FnCondition,
		parse.FnEquals,
		parse.FnFindInMap,
		parse.FnIf,
		parse.FnNot,
		parse.FnOr,
		parse.FnRef,
	}
	for _, fn := range validFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnNot)(ExampleValidIFs[fn]()), ValueString, true, fmt.Sprintf("%s allowed as condition", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnNot)(ExampleValidIFs[fn]()), ValueString, false, fmt.Sprintf("%s not allowed as condition", fn)})
	}

	scenarios.evaluate(t, validateNot, ctx)
}
