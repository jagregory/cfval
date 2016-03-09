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

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnCondition)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnCondition)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Condition", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Condition", map[string]interface{}{"Condition": "Condition", "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnCondition)("NotACondition"), ValueString, false, "invalid condition"},
		IFScenario{IF(parse.FnCondition)("Condition"), ValueString, true, "valid condition"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnCondition)(ExampleValidIFs[fn]()), ValueString, false, fmt.Sprintf("%s not allowed as condition name", fn)})
	}

	scenarios.evaluate(t, validateCondition, ctx)
}
