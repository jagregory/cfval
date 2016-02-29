package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

type IFScenario struct {
	fn      parse.IntrinsicFunction
	pass    bool
	message string
}

func TestAnd(t *testing.T) {
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

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnAnd)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnAnd)(nil), false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{"Fn::And": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), false, "too many arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition]()}), false, "too few arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), true, "minimum arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), true, "some arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), true, "maximum arguments"},
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
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[fn](), ExampleValidIFs[parse.FnCondition]()}), true, fmt.Sprintf("%s allowed as condition", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[fn]()}), true, fmt.Sprintf("%s allowed as condition", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[fn](), ExampleValidIFs[parse.FnCondition]()}), false, fmt.Sprintf("%s not allowed as condition", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s not allowed as condition", fn)})
	}

	for i, s := range scenarios {
		errs := validateAnd(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
