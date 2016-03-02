package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJoin(t *testing.T) {
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
		IFScenario{IF(parse.FnJoin)(float64(123)), false, "invalid type used for arg"},
		IFScenario{IF(parse.FnJoin)(nil), false, "nil used for arg"},
		IFScenario{IF(parse.FnJoin)([]interface{}{}), false, "no args"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", "b", "c"}), false, "too many args"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a"}), false, "too few args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{"Fn::Join": []interface{}{"delim", []interface{}{"a", "b"}}, "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnJoin)([]interface{}{float64(1), []interface{}{"b", "c"}}), false, "invalid type for delimiter"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", float64(1)}), false, "invalid type for values"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", []interface{}{"a", float64(1)}}), false, "invalid type used in values"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", []interface{}{"a", "b"}}), true, "valid"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{ExampleValidIFs[fn](), []interface{}{"a", "b"}}), false, fmt.Sprintf("%s as delimiter", fn)})
	}

	validValuesFns := []parse.IntrinsicFunctionSignature{
		parse.FnBase64,
		parse.FnFindInMap,
		parse.FnGetAtt,
		parse.FnGetAZs,
		parse.FnIf,
		parse.FnJoin,
		parse.FnSelect,
		parse.FnRef,
	}
	for _, fn := range validValuesFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), true, fmt.Sprintf("%s is allowed as values", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), true, fmt.Sprintf("%s is allowed as a value", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validValuesFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s is not allowed as values", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), false, fmt.Sprintf("%s is not allowed as a value", fn)})
	}

	for i, s := range scenarios {
		errs := validateJoin(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
