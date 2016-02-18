package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestSingleValueValidate(t *testing.T) {
	type testcase struct{ expected, actual interface{} }
	pass := []testcase{
		testcase{expected: float64(10), actual: float64(10)},
		testcase{expected: "hi", actual: "hi"},
	}
	fail := []testcase{
		testcase{expected: float64(10), actual: float64(9)},
		testcase{expected: float64(10), actual: float64(11)},
		testcase{expected: float64(10), actual: "10"},
		testcase{expected: "hi", actual: "10"},
		testcase{expected: "hi", actual: float64(10)},
	}

	prop := Schema{}
	template := &parse.Template{}
	tr := parse.TemplateResource{
		Tmpl: template,
	}
	currentResource := ResourceWithDefinition{tr, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	for _, test := range pass {
		if _, errs := SingleValueValidate(test.expected)(prop, test.actual, currentResource, ctx); errs != nil {
			t.Errorf("Should pass with expected %s and actual %s", test.expected, test.actual)
		}
	}

	for _, test := range fail {
		if _, errs := SingleValueValidate(test.expected)(prop, test.actual, currentResource, ctx); errs == nil {
			t.Errorf("Should fail with expected %s and actual %s", test.expected, test.actual)
		}
	}
}

func TestFixedArrayValidateHelper(t *testing.T) {
	template := &parse.Template{}
	tr := parse.TemplateResource{
		Tmpl: template,
	}
	currentResource := ResourceWithDefinition{tr, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	validate := FixedArrayValidate([]string{"a", "b", "c"}, []string{"d", "e"})

	if _, errs := validate(Schema{}, []interface{}{}, currentResource, ctx); errs == nil {
		t.Error("Should fail on empty list")
	}

	if _, errs := validate(Schema{}, []interface{}{"c", "d"}, currentResource, ctx); errs == nil {
		t.Error("Should fail on unexpected list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "b"}, currentResource, ctx); errs == nil {
		t.Error("Should fail on subset list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "b", "c"}, currentResource, ctx); errs != nil {
		t.Error("Should pass on expected list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "c", "b"}, currentResource, ctx); errs != nil {
		t.Error("Should pass on unordered expected list")
	}
}

func TestRegexpValidateHelper(t *testing.T) {
	template := &parse.Template{}
	validator := RegexpValidate("^a string$", "Match failed")
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	_, errs := validator(Schema{}, "a string", currentResource, ctx)
	if errs != nil {
		t.Error("Should pass on a valid string")
	}
	if len(errs) > 0 {
		t.Error("Should pass on a valid string", errs)
	}

	_, errs = validator(Schema{}, "no match", currentResource, ctx)
	if errs == nil {
		t.Error("Should fail on a non-matching string")
	}

	if len(errs) == 0 || errs[0].Message != "Match failed" {
		t.Error("Should fail with supplied message")
	}
}

func TestIntegerRangeValidateHelper(t *testing.T) {
	template := &parse.Template{}
	validator := IntegerRangeValidate(5, 10)
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	for _, valid := range []float64{5, 6, 7, 8, 9, 10} {
		if _, errs := validator(Schema{}, valid, currentResource, ctx); errs != nil {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []float64{-10, 0, 1, 2, 11, 100} {
		if _, errs := validator(Schema{}, invalid, currentResource, ctx); errs == nil {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}

func TestStringLengthValidateHelper(t *testing.T) {
	template := &parse.Template{}
	validator := StringLengthValidate(5, 10)
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	for _, valid := range []string{"abcde", "abcdefghij"} {
		if _, errs := validator(Schema{}, valid, currentResource, ctx); errs != nil {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []string{"", "abcd", "abcdefghijk"} {
		if _, errs := validator(Schema{}, invalid, currentResource, ctx); errs == nil {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}

func TestNumberOptionsValidateHelper(t *testing.T) {
	template := &parse.Template{}
	validator := NumberOptions(5, 10, 20)
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	for _, valid := range []float64{5, 10, 20} {
		if _, errs := validator(Schema{}, valid, currentResource, ctx); errs != nil {
			t.Error("Should pass on valid value", valid, errs)
		}
	}

	for _, invalid := range []float64{-10, 0, 6, 1000} {
		if _, errs := validator(Schema{}, invalid, currentResource, ctx); errs == nil {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}
