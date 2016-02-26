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

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnIf)(123), false, "invalid type used for arg"},
		IFScenario{IF(parse.FnIf)(nil), false, "nil used for arg"},
		IFScenario{IF(parse.FnIf)([]interface{}{}), false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::If", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::If", map[string]interface{}{"Fn::If": "", "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnIf)([]interface{}{"NotACondition", "a", "b"}), false, "not a valid condition name"},
		IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", "b"}), true, "valid condition name"},
		IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", IF(parse.FnRef)("AWS::NoValue")}), true, "AWS::NoValue"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{ExampleValidIFs[fn](), "a", "b"}), false, fmt.Sprintf("%s in place of condition name", fn)})
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
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", ExampleValidIFs[fn](), "b"}), true, fmt.Sprintf("%s allowed for true value", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", ExampleValidIFs[fn]()}), true, fmt.Sprintf("%s allowed for false value", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validValueFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", ExampleValidIFs[fn](), "b"}), false, fmt.Sprintf("%s not allowed for true value", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnIf)([]interface{}{"Condition", "a", ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s not allowed for false value", fn)})
	}

	for i, s := range scenarios {
		_, errs := validateIf(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
