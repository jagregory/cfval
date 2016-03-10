package schema

import (
	"fmt"
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJoinDelimiter(t *testing.T) {
	template := &parse.Template{
		Conditions: map[string]parse.Condition{
			"Condition": parse.Condition{},
		},

		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},

			"MyListResource": parse.TemplateResource{
				Type: "ListResource",
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

		"ListResource": Resource{
			ReturnValue: Schema{
				Type: Multiple(ValueString),
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnJoin)(float64(123)), ValueString, false, "invalid type used for arg"},
		IFScenario{IF(parse.FnJoin)(nil), ValueString, false, "nil used for arg"},
		IFScenario{IF(parse.FnJoin)([]interface{}{}), ValueString, false, "no args"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", "b", "c"}), ValueString, false, "too many args"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"a"}), ValueString, false, "too few args"},
		IFScenario{parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{}}, ValueString, false, "empty map"},
		IFScenario{parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{"Fn::Join": []interface{}{"delim", []interface{}{"a", "b"}}, "blah": "blah"}}, ValueString, false, "extra properties"},
		IFScenario{IF(parse.FnJoin)([]interface{}{float64(1), []interface{}{"b", "c"}}), ValueString, false, "invalid type for delimiter"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", []interface{}{"a", "b"}}), ValueString, true, "valid"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{ExampleValidIFs[fn](), []interface{}{"a", "b"}}), ValueString, false, fmt.Sprintf("%s as delimiter", fn)})
	}

	scenarios.evaluate(t, validateJoin, ctx)
}

func TestJoinItems(t *testing.T) {
	template := &parse.Template{
		Conditions: map[string]parse.Condition{
			"Condition": parse.Condition{},
		},

		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},

			"MyListResource": parse.TemplateResource{
				Type: "ListResource",
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

		"ListResource": Resource{
			ReturnValue: Schema{
				Type: Multiple(ValueString),
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", []interface{}{"a", float64(1)}}), ValueString, false, "invalid type used in values"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", []interface{}{"a", "b"}}), ValueString, true, "valid"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", []interface{}{IF(parse.FnRef)("MyResource"), "b"}}), ValueBool, true, "nested-Ref shouldn't care about PropertyType"},
	}

	validValuesFns := []parse.IntrinsicFunctionSignature{
		parse.FnBase64,
		parse.FnFindInMap,
		parse.FnGetAtt,
		parse.FnGetAZs,
		parse.FnIf,
		parse.FnJoin,
		parse.FnSelect,
		parse.FnRef,
	}
	for _, fn := range validValuesFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), ValueString, true, fmt.Sprintf("%s is allowed as a value", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validValuesFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), ValueString, false, fmt.Sprintf("%s is not allowed as a value", fn)})
	}

	scenarios.evaluate(t, validateJoin, ctx)
}

func TestJoinItemArray(t *testing.T) {
	template := &parse.Template{
		Conditions: map[string]parse.Condition{
			"Condition": parse.Condition{},
		},

		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},

			"MyListResource": parse.TemplateResource{
				Type: "ListResource",
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
					Type: Multiple(ValueString),
				},
			},

			ReturnValue: Schema{
				Type: Multiple(ValueString),
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	scenarios := IFScenarios{
		IFScenario{IF(parse.FnJoin)([]interface{}{"a", float64(1)}), ValueString, false, "invalid type for values"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", []interface{}{"a", "b"}}), ValueString, true, "valid"},
		IFScenario{IF(parse.FnJoin)([]interface{}{"d", IF(parse.FnRef)("MyResource")}), ValueBool, true, "nested-Ref shouldn't care about PropertyType"},
	}

	for _, fn := range parse.AllIntrinsicFunctions {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{ExampleValidIFs[fn](), []interface{}{"a", "b"}}), ValueString, false, fmt.Sprintf("%s as delimiter", fn)})
	}

	validValuesFns := []parse.IntrinsicFunctionSignature{
		parse.FnBase64,
		parse.FnFindInMap,
		parse.FnGetAtt,
		parse.FnGetAZs,
		parse.FnIf,
		parse.FnJoin,
		parse.FnSelect,
		parse.FnRef,
	}
	for _, fn := range validValuesFns {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), ValueString, true, fmt.Sprintf("%s is allowed as values", fn)})
	}
	for _, fn := range parse.AllIntrinsicFunctions.Except(validValuesFns...) {
		scenarios = append(scenarios, IFScenario{IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), ValueString, false, fmt.Sprintf("%s is not allowed as values", fn)})
	}

	scenarios.evaluate(t, validateJoin, ctx)
}
