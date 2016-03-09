package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestBase64(t *testing.T) {
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
		IFScenario{IF(parse.FnBase64)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnBase64)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Base64", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Base64", map[string]interface{}{"Fn::Base64": "example", "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnBase64)("example"), ValueString, true, "valid value used"},
		IFScenario{IF(parse.FnBase64)(ExampleValidIFs[parse.FnIf]()), ValueString, true, "If used"},
	}

	validIfs := []parse.IntrinsicFunctionSignature{
		parse.FnIf,
		parse.FnJoin,
		parse.FnRef,
	}
	for _, fn := range validIfs {
		scenarios = append(scenarios, IFScenario{IF(parse.FnBase64)(ExampleValidIFs[fn]()), ValueString, true, fmt.Sprintf("%s as value", fn)})
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(validIfs...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnBase64)(ExampleValidIFs[fn]()), ValueString, false, fmt.Sprintf("%s as value", fn)})
	}

	scenarios.evaluate(t, validateBase64, ctx)
}
