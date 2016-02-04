package schema

import "testing"

func TestRefValidate(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{},
		},
		Parameters: map[string]Parameter{
			"MyParameter": Parameter{},
		},
	}
	resourceRef := NewRef("MyResource")
	parameterRef := NewRef("MyParameter")
	pseudoParameterRef := NewRef("AWS::StackName")
	invalidRef := NewRef("invalid")

	if ok, _ := resourceRef.Validate(template, []string{}); !ok {
		t.Error("Should pass on valid resource ref", resourceRef.target)
	}

	if ok, _ := parameterRef.Validate(template, []string{}); !ok {
		t.Error("Should pass on valid parameter ref", parameterRef.target)
	}

	if ok, _ := pseudoParameterRef.Validate(template, []string{}); !ok {
		t.Error("Should pass on valid pseudo parameter ref", pseudoParameterRef.target)
	}

	if ok, _ := invalidRef.Validate(template, []string{}); ok {
		t.Error("Should fail on invalid ref", invalidRef.target)
	}
}
