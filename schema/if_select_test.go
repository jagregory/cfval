package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSelectIndex(t *testing.T) {
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
				"Name": Schema{
					Type: ValueNumber,
				},
			},

			ReturnValue: Schema{
				Type: ValueNumber,
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
		IFScenario{IF(parse.FnSelect)([]interface{}{IF(parse.FnRef)("MyResource"), validArray}), ValueBool, true, "Nested-FN works with PropertyType in index"},
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

	scenarios.evaluate(t, validateSelect, ctx)
}

func TestSelectArray(t *testing.T) {
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
				"Name": Schema{
					Type: Multiple(ValueString),
				},
			},

			ReturnValue: Schema{
				Type: Multiple(ValueString),
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), nil}), ValueString, false, "nil array"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), []interface{}{"one", nil, "three"}}), ValueString, false, "array has nils"},
		IFScenario{IF(parse.FnSelect)([]interface{}{float64(1), IF(parse.FnGetAtt)([]interface{}{"MyResource", "Name"})}), ValueBool, true, "Nested-FN works with PropertyType in index"},
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
