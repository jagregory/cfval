package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestCondition(t *testing.T) {
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
		IFScenario{IF(parse.FnCondition)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnCondition)(nil), false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Condition", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Condition", map[string]interface{}{"Condition": "Condition", "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnCondition)("NotACondition"), false, "invalid condition"},
		IFScenario{IF(parse.FnCondition)("Condition"), true, "valid condition"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnCondition)(ExampleValidIFs[fn]()), false, fmt.Sprintf("%s not allowed as condition name", fn)})
	}

	for i, s := range scenarios {
		_, errs := validateCondition(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
