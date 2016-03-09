package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestGetAtt(t *testing.T) {
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
		IFScenario{IF(parse.FnGetAtt)(123), InstanceID, false, "invalid type used for args"},
		IFScenario{IF(parse.FnGetAtt)(nil), InstanceID, false, "nil used for args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{}), InstanceID, false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{}}, InstanceID, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"example", "example"}, "blah": "blah"}}, InstanceID, false, "extra properties"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"a", "b", "c"}), InstanceID, false, "too many args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"a"}), InstanceID, false, "too few args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"UnknownResource", "prop"}), InstanceID, false, "invalid resource"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "BadProp"}), InstanceID, false, "invalid property used for type of resource"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "Name"}), InstanceID, false, "valid property of wrong type"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "InstanceId"}), InstanceID, true, "valid property"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", ExampleValidIFs[parse.FnRef]()}), InstanceID, true, "Ref in Attribute"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "ListInstanceId"}), Multiple(InstanceID), true, "valid property used for type of resource"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAtt)([]interface{}{ExampleValidIFs[fn](), "InstanceId"}), InstanceID, false, fmt.Sprintf("%s as Resource", fn)})
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", ExampleValidIFs[fn]()}), InstanceID, false, fmt.Sprintf("%s in Attribute", fn)})
	}

	scenarios.evaluate(t, validateGetAtt, ctx)
}
