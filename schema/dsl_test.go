package schema

import "testing"

func TestFixedArrayValidateHelper(t *testing.T) {
	template := &Template{}
	tr := TemplateResource{
		Template: template,
	}
	context := []string{}

	validate := FixedArrayValidate([]string{"a", "b", "c"}, []string{"d", "e"})

	if ok, _ := validate([]interface{}{}, tr, context); ok {
		t.Error("Should fail on empty list")
	}

	if ok, _ := validate([]interface{}{"c", "d"}, tr, context); ok {
		t.Error("Should fail on unexpected list")
	}

	if ok, _ := validate([]interface{}{"a", "b"}, tr, context); ok {
		t.Error("Should fail on subset list")
	}

	if ok, _ := validate([]interface{}{"a", "b", "c"}, tr, context); !ok {
		t.Error("Should pass on expected list")
	}

	if ok, _ := validate([]interface{}{"a", "c", "b"}, tr, context); !ok {
		t.Error("Should pass on unordered expected list")
	}
}

func TestRegexpValidateHelper(t *testing.T) {
	validator := RegexpValidate("^a string$", "Match failed")

	ok, errs := validator("a string", TemplateResource{}, []string{})
	if !ok {
		t.Error("Should pass on a valid string")
	}
	if len(errs) > 0 {
		t.Error("Should pass on a valid string", errs)
	}

	ok, errs = validator("no match", TemplateResource{}, []string{})
	if ok {
		t.Error("Should fail on a non-matching string")
	}

	if len(errs) == 0 || errs[0].Message != "Match failed" {
		t.Error("Should fail with supplied message")
	}
}

func TestIntegerRangeValidateHelper(t *testing.T) {
	validator := IntegerRangeValidate(5, 10)

	for _, valid := range []float64{5, 6, 7, 8, 9, 10} {
		if ok, _ := validator(valid, TemplateResource{}, []string{}); !ok {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []float64{-10, 0, 1, 2, 11, 100} {
		if ok, _ := validator(invalid, TemplateResource{}, []string{}); ok {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}

func TestStringLengthValidateHelper(t *testing.T) {
	validator := StringLengthValidate(5, 10)

	for _, valid := range []string{"abcde", "abcdefghij"} {
		if ok, _ := validator(valid, TemplateResource{}, []string{}); !ok {
			t.Error("Should pass on valid value", valid)
		}
	}

	for _, invalid := range []string{"", "abcd", "abcdefghijk"} {
		if ok, _ := validator(invalid, TemplateResource{}, []string{}); ok {
			t.Error("Should fail on invalid value", invalid)
		}
	}
}
