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
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnBase64)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnBase64)(nil), false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Base64", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Base64", map[string]interface{}{"Fn::Base64": "example", "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnBase64)("example"), true, "valid value used"},
		IFScenario{IF(parse.FnBase64)(ExampleValidIFs[parse.FnIf]()), true, "If used"},
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnIf) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnBase64)(ExampleValidIFs[fn]()), false, fmt.Sprintf("%s as value", fn)})
	}

	for i, s := range scenarios {
		errs := validateBase64(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
