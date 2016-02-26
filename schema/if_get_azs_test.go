package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestGetAZs(t *testing.T) {
	template := &parse.Template{
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
		IFScenario{IF(parse.FnGetAZs)(123), false, "invalid type used for arg"},
		IFScenario{IF(parse.FnGetAZs)(nil), false, "nil used for arg"},
		IFScenario{IF(parse.FnGetAZs)([]interface{}{}), false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": "", "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnGetAZs)(""), true, "empty arg"},
		IFScenario{IF(parse.FnGetAZs)("ap-southeast-2"), true, "valid region"},
		// TODO: IFScenario{IF(parse.FnGetAZs)("ap-southeast-9"), false, "invalid region"},
		IFScenario{IF(parse.FnGetAZs)(ExampleValidIFs[parse.FnRef]()), true, "Ref used as arg"},
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAZs)(ExampleValidIFs[fn]()), false, fmt.Sprintf("%s as arg", fn)})
	}

	for i, s := range scenarios {
		_, errs := validateGetAZs(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
