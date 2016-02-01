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
