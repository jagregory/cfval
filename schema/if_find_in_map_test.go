package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestFindInMap(t *testing.T) {
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
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := []IFScenario{
		IFScenario{IF(parse.FnFindInMap)(123), false, "invalid type used for args"},
		IFScenario{IF(parse.FnFindInMap)(nil), false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{}}, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": "example", "blah": "blah"}}, false, "extra properties"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a", "b", "c", "d"}), false, "too many args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a"}), false, "too few args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a", 1, "c"}), false, "invalid type in args"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", "subkey"}), true, "valid args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[parse.FnRef](), "key", "subkey"}), true, "Ref as MapName"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[parse.FnRef](), "subkey"}), true, "Ref as TopLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[parse.FnRef]()}), true, "Ref as SecondLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[parse.FnFindInMap](), "key", "subkey"}), true, "FindInMap as MapName"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[parse.FnFindInMap](), "subkey"}), true, "FindInMap as TopLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[parse.FnFindInMap]()}), true, "FindInMap as SecondLevelKey"},
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef, parse.FnFindInMap) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[fn](), "key", "subkey"}), false, fmt.Sprintf("%s as MapName", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[fn](), "subkey"}), false, fmt.Sprintf("%s as TopLevelKey", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[fn]()}), false, fmt.Sprintf("%s as SecondLevelKey", fn)})
	}

	for i, s := range scenarios {
		errs := validateFindInMap(s.fn, ctx)
		if s.pass && errs != nil {
			t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
		} else if !s.pass && errs == nil {
			t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
		}
	}
}
