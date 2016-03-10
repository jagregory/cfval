package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSelect(t *testing.T) {
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

				"ListInstanceId": Schema{
					Type: Multiple(InstanceID),
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

	validArray := []interface{}{"a", "b", "c"}
	scenarios := IFScenarios{
		IFScenario{IF(parse.FnSelect)(123), ValueString, false, "invalid type used for args"},
		IFScenario{IF(parse.FnSelect)(nil), ValueString, false, "nil used for args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Select", map[string]interface{}{"Fn::Select": []interface{}{float64(1), validArray}, "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), validArray}), ValueString, true, "valid index and array"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(10), validArray}), ValueString, false, "index out of bounds"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(-1), validArray}), ValueString, false, "negative index"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), nil}), ValueString, false, "nil array"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), []interface{}{"one", nil, "three"}}), ValueString, false, "array has nils"},
	}

	validIndexFns := []parse.IntrinsicFunctionSignature{
		parse.FnFindInMap,
		parse.FnRef,
	}
	for _, fn := range validIndexFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnSelect)([]interface{}{ExampleValidIFs[fn](), validArray}), ValueString, true, fmt.Sprintf("%s allowed as index", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validIndexFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnSelect)([]interface{}{ExampleValidIFs[fn](), validArray}), ValueString, false, fmt.Sprintf("%s not allowed as index", fn)})
	}

	validArrayFns := []parse.IntrinsicFunctionSignature{
		parse.FnFindInMap,
		parse.FnRef,
		parse.FnGetAtt,
		parse.FnGetAZs,
		parse.FnIf,
	}
	for _, fn := range validArrayFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[fn]()}), ValueString, true, fmt.Sprintf("%s allowed as array", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validArrayFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), ExampleValidIFs[fn]()}), ValueString, false, fmt.Sprintf("%s not allowed as array", fn)})
	}

	scenarios.evaluate(t, validateSelect, ctx)
}
