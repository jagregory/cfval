package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestIf(t *testing.T) {
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

				"ListInstanceId": Schema{
					Type: Multiple(InstanceID),
				},

				"Name": Schema{
					Type: ValueString,
				},
			},

			ReturnValue: Schema{
				Type: InstanceID,
			},
		},
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnIf)(123), InstanceID, false, "invalid type used for arg"},
		IFScenario{IF(parse.FnIf)(nil), InstanceID, false, "nil used for arg"},
		IFScenario{IF(parse.FnIf)([]interface{}{}), InstanceID, false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::If", map[string]interface{}{}}, InstanceID, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::If", map[string]interface{}{"Fn::If": "", "blah": "blah"}}, InstanceID, false, "extra properties"},
		IFScenario{IF(parse.FnIf)([]interface{}{"NotACondition", "a", "b"}), InstanceID, false, "not a valid condition name"},
		IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", "b"}), InstanceID, true, "valid condition name"},
		IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", IF(parse.FnRef)("AWS::NoValue")}), InstanceID, true, "AWS::NoValue"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{ExampleValidIFs[fn](), "a", "b"}), InstanceID, false, fmt.Sprintf("%s in place of condition name", fn)})
	}

	validValueFns := []parse.IntrinsicFunctionSignature{
		parse.FnAnd,
		parse.FnCondition,
		parse.FnEquals,
		parse.FnFindInMap,
		parse.FnIf,
		parse.FnNot,
		parse.FnOr,
		parse.FnRef,
	}
	for _, fn := range validValueFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", ExampleValidIFs[fn](), "b"}), InstanceID, true, fmt.Sprintf("%s allowed for true value", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", ExampleValidIFs[fn]()}), InstanceID, true, fmt.Sprintf("%s allowed for false value", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validValueFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", ExampleValidIFs[fn](), "b"}), InstanceID, false, fmt.Sprintf("%s not allowed for true value", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", ExampleValidIFs[fn]()}), InstanceID, false, fmt.Sprintf("%s not allowed for false value", fn)})
	}

	scenarios.evaluate(t, validateIf, ctx)
}
