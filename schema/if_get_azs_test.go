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

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnGetAZs)(123), InstanceID, false, "invalid type used for arg"},
		IFScenario{IF(parse.FnGetAZs)(nil), InstanceID, false, "nil used for arg"},
		IFScenario{IF(parse.FnGetAZs)([]interface{}{}), InstanceID, false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{}}, InstanceID, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAZs", map[string]interface{}{"Fn::GetAZs": "", "blah": "blah"}}, InstanceID, false, "extra properties"},
		IFScenario{IF(parse.FnGetAZs)(""), InstanceID, true, "empty arg"},
		IFScenario{IF(parse.FnGetAZs)("ap-southeast-2"), InstanceID, true, "valid region"},
		// TODO: IFScenario{IF(parse.FnGetAZs)("ap-southeast-9"), InstanceID, false, "invalid region"},
		IFScenario{IF(parse.FnGetAZs)(ExampleValidIFs[parse.FnRef]()), InstanceID, true, "Ref used as arg"},
		IFScenario{IF(parse.FnGetAZs)(ExampleValidIFs[parse.FnRef]()), ValueBool, true, "Ref used as arg with different property type"},
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAZs)(ExampleValidIFs[fn]()), InstanceID, false, fmt.Sprintf("%s as arg", fn)})
	}

	scenarios.evaluate(t, validateGetAZs, ctx)
}
