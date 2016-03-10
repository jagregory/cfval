package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

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

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnAnd)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnAnd)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{"Fn::And": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), ValueString, false, "too many arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition]()}), ValueString, false, "too few arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), ValueString, true, "minimum arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{true, true}), ValueString, false, "invalid arguments (booleans not allowed)"},
		IFScenario{IF(parse.FnAnd)([]interface{}{"true", "true"}), ValueString, false, "invalid arguments (strings not allowed)"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), ValueString, true, "some arguments"},
		IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[parse.FnCondition]()}), ValueString, true, "maximum arguments"},
	}

	validFns := []parse.IntrinsicFunctionSignature{
		parse.FnAnd,
		parse.FnCondition,
		parse.FnEquals,
		parse.FnIf,
		parse.FnNot,
		parse.FnOr,
	}
	for _, fn := range validFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[fn](), ExampleValidIFs[parse.FnCondition]()}), ValueString, true, fmt.Sprintf("%s allowed as condition", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[fn]()}), ValueString, true, fmt.Sprintf("%s allowed as condition", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[fn](), ExampleValidIFs[parse.FnCondition]()}), ValueString, false, fmt.Sprintf("%s not allowed as condition", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnAnd)([]interface{}{ExampleValidIFs[parse.FnCondition](), ExampleValidIFs[fn]()}), ValueString, false, fmt.Sprintf("%s not allowed as condition", fn)})
	}

	scenarios.evaluate(t, validateAnd, ctx)
}
