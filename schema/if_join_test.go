package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestJoin(t *testing.T) {
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

				"Name": Schema{
					Type: ValueString,
				},
			},

			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: ValueString}, ValidationOptions{})

	if _, errs := validateJoin(IF(parse.FnJoin)(123), ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateJoin(parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when no args", errs)
	}

	if _, errs := validateJoin(parse.IntrinsicFunction{"Fn::Join", map[string]interface{}{"Fn::Join": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when valid with extra properties", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"a", "b", "c"}), ctx); errs == nil {
		t.Error("Should fail when too many arguments supplied", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"a"}), ctx); errs == nil {
		t.Error("Should fail when too few arguments supplied", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{1, []interface{}{"a", "b"}}), ctx); errs == nil {
		t.Error("Should fail when invalid type used for delimeter", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"a", 1}), ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"a", []interface{}{1, "b"}}), ctx); errs == nil {
		t.Error("Should fail when invalid type used in args", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"d", []interface{}{"a", "b"}}), ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"d", []interface{}{ExampleValidIFs[parse.FnRef](), "b", "c"}}), ctx); errs != nil {
		t.Error("Should short circuit and pass when ref used", errs)
	}

	invalidDelimeterFns := parse.AllIntrinsicFunctions
	for _, fn := range invalidDelimeterFns {
		if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{ExampleValidIFs[fn](), []interface{}{"a", "b"}}), ctx); errs == nil {
			t.Errorf("Should fail with %s in Delimeter: %s", fn, errs)
		}
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
		if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), ctx); errs != nil {
			t.Errorf("%s is allowed as values (errs: %s)", fn, errs)
		}

		if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), ctx); errs != nil {
			t.Errorf("%s is allowed as a value (errs: %s)", fn, errs)
		}
	}

	for _, fn := range parse.AllIntrinsicFunctions.Except(validValuesFns...) {
		if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"delim", ExampleValidIFs[fn]()}), ctx); errs == nil {
			t.Errorf("Should fail with %s for values: %s", fn, errs)
		}

		if _, errs := validateJoin(IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"a", ExampleValidIFs[fn]()}}), ctx); errs == nil {
			t.Errorf("Should fail with %s for a value: %s", fn, errs)
		}
	}
}
