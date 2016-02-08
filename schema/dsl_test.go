package schema

import "testing"

func TestFixedArrayValidateHelper(t *testing.T) {
	tr := TemplateResource{
		template: &Template{},
	}
	context := []string{}

	validate := FixedArrayValidate([]string{"a", "b", "c"}, []string{"d", "e"})

	if _, errs := validate(Schema{}, []interface{}{}, tr, context); errs == nil {
		t.Error("Should fail on empty list")
	}

	if _, errs := validate(Schema{}, []interface{}{"c", "d"}, tr, context); errs == nil {
		t.Error("Should fail on unexpected list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "b"}, tr, context); errs == nil {
		t.Error("Should fail on subset list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "b", "c"}, tr, context); errs != nil {
		t.Error("Should pass on expected list")
	}

	if _, errs := validate(Schema{}, []interface{}{"a", "c", "b"}, tr, context); errs != nil {
		t.Error("Should pass on unordered expected list")
	}
}

func TestRegexpValidateHelper(t *testing.T) {
	validator := RegexpValidate("^a string$", "Match failed")

	_, errs := validator(Schema{}, "a string", TemplateResource{}, []string{})
	if errs != nil {
		t.Error("Should pass on a valid string")
	}
	if len(errs) > 0 {
		t.Error("Should pass on a valid string", errs)
	}

	_, errs = validator(Schema{}, "no match", TemplateResource{}, []string{})
	if errs == nil {
		t.Error("Should fail on a non-matching string")
	}

	if len(errs) == 0 || errs[0].Message != "Match failed" {
		t.Error("Should fail with supplied message")
	}
}

func TestIntegerRangeValidateHelper(t *testing.T) {
	validator := IntegerRangeValidate(5, 10)

	for _, valid := range []float64{5, 6, 7, 8, 9, 10} {
		if _, errs := validator(Schema{}, valid, TemplateResource{}, []string{}); errs != nil {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []float64{-10, 0, 1, 2, 11, 100} {
		if _, errs := validator(Schema{}, invalid, TemplateResource{}, []string{}); errs == nil {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}

func TestStringLengthValidateHelper(t *testing.T) {
	validator := StringLengthValidate(5, 10)

	for _, valid := range []string{"abcde", "abcdefghij"} {
		if _, errs := validator(Schema{}, valid, TemplateResource{}, []string{}); errs != nil {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []string{"", "abcd", "abcdefghijk"} {
		if _, errs := validator(Schema{}, invalid, TemplateResource{}, []string{}); errs == nil {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}
