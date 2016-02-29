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

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnGetAtt)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnGetAtt)(nil), false, "nil used for args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{}), false, "no args"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::GetAtt", map[string]interface{}{"Fn::GetAtt": []interface{}{"example", "example"}, "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"a", "b", "c"}), false, "too many args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"a"}), false, "too few args"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"UnknownResource", "prop"}), false, "invalid resource"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "BadProp"}), false, "invalid property used for type of resource"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "Name"}), false, "valid property of wrong type"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", "InstanceId"}), true, "valid property"},
		IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", ExampleValidIFs[parse.FnRef]()}), true, "Ref in Attribute"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAtt)([]interface{}{ExampleValidIFs[fn](), "InstanceId"}), false, fmt.Sprintf("%s as Resource", fn)})
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnGetAtt)([]interface{}{"MyResource", ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s in Attribute", fn)})
	}

	for i, s := range scenarios {
		errs := validateGetAtt(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}

	listCtx := NewPropertyContext(ctx, Schema{Type: Multiple(InstanceID)})
	if errs := validateGetAtt(IF(parse.FnGetAtt)([]interface{}{"MyResource", "ListInstanceId"}), listCtx); errs != nil {
		t.Error("Should pass when valid property used for type of resource", errs)
	}
}
