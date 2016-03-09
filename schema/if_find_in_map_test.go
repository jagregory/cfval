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

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnFindInMap)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnFindInMap)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::FindInMap", map[string]interface{}{"Fn::FindInMap": "example", "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a", "b", "c", "d"}), ValueString, false, "too many args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a"}), ValueString, false, "too few args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"a", 1, "c"}), ValueString, false, "invalid type in args"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", "subkey"}), ValueString, true, "valid args used"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[parse.FnRef](), "key", "subkey"}), ValueString, true, "Ref as MapName"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[parse.FnRef](), "subkey"}), ValueString, true, "Ref as TopLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[parse.FnRef]()}), ValueString, true, "Ref as SecondLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[parse.FnFindInMap](), "key", "subkey"}), ValueString, true, "FindInMap as MapName"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[parse.FnFindInMap](), "subkey"}), ValueString, true, "FindInMap as TopLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[parse.FnFindInMap]()}), ValueString, true, "FindInMap as SecondLevelKey"},
		IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[parse.FnRef](), ExampleValidIFs[parse.FnRef](), ExampleValidIFs[parse.FnRef]()}), InstanceID, true, "nested-Ref able to be coerced"},
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(parse.FnRef, parse.FnFindInMap) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{ExampleValidIFs[fn](), "key", "subkey"}), ValueString, false, fmt.Sprintf("%s as MapName", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", ExampleValidIFs[fn](), "subkey"}), ValueString, false, fmt.Sprintf("%s as TopLevelKey", fn)})
		scenarios = append(scenarios, IFScenario{IF(parse.FnFindInMap)([]interface{}{"map", "key", ExampleValidIFs[fn]()}), ValueString, false, fmt.Sprintf("%s as SecondLevelKey", fn)})
	}

	scenarios.evaluate(t, validateFindInMap, ctx)
}
