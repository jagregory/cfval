package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestEqual(t *testing.T) {
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
		IFScenario{IF(parse.FnEquals)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnEquals)(nil), false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Equals", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Equals", map[string]interface{}{"Fn::Equals": []interface{}{"a", "b"}, "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnEquals)([]interface{}{"a", "b", "c"}), false, "too many arguments"},
		IFScenario{IF(parse.FnEquals)([]interface{}{"a"}), false, "too few arguments"},
		IFScenario{IF(parse.FnEquals)([]interface{}{nil, "a"}), false, "nil in left"},
		IFScenario{IF(parse.FnEquals)([]interface{}{"a", nil}), false, "nil in right"},
		IFScenario{IF(parse.FnEquals)([]interface{}{"a", "b"}), true, "both strings"},
		IFScenario{IF(parse.FnEquals)([]interface{}{"a", 1}), true, "mixed types"},
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
		scenarios = append(scenarios, IFScenario{IF(parse.FnEquals)([]interface{}{ExampleValidIFs[fn](), "b"}), true, fmt.Sprintf("%s allowed as left", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnEquals)([]interface{}{"a", ExampleValidIFs[fn]()}), true, fmt.Sprintf("%s allowed as right", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnEquals)([]interface{}{ExampleValidIFs[fn](), "b"}), false, fmt.Sprintf("%s not allowed as left", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnEquals)([]interface{}{"a", ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s not allowed as right", fn)})
	}

	for i, s := range scenarios {
		errs := validateEquals(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
