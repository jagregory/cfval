package schema

import "testing"

func TestRefValidate(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"Resource1": TemplateResource{},
			"Resource2": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: TypeString,
					},
				},
			},
		},
		Parameters: map[string]Parameter{
			"Parameter1": Parameter{},
			"Parameter2": Parameter{
				Type: TypeString,
			},
		},
	}

	if ok, _ := NewRef(Schema{Type: TypeString}, "Resource1").Validate(template, []string{}); !ok {
		t.Error("Should pass on valid resource ref with Unknown ref type")
	}

	if ok, _ := NewRef(Schema{Type: TypeString}, "Resource2").Validate(template, []string{}); !ok {
		t.Error("Should pass on valid resource ref with matching types")
	}

	if ok, _ := NewRef(Schema{Type: TypeInteger}, "Resource2").Validate(template, []string{}); ok {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if ok, _ := NewRef(Schema{Type: TypeString}, "Parameter1").Validate(template, []string{}); !ok {
		t.Error("Should pass on valid parameter ref with Unknown ref type")
	}

	if ok, _ := NewRef(Schema{Type: TypeString}, "Parameter2").Validate(template, []string{}); !ok {
		t.Error("Should pass on valid parameter ref with matching types")
	}

	if ok, _ := NewRef(Schema{Type: TypeInteger}, "Parameter2").Validate(template, []string{}); ok {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if ok, _ := NewRef(Schema{Type: TypeString}, "AWS::StackName").Validate(template, []string{}); !ok {
		t.Error("Should pass on valid pseudo-parameter ref with matching types")
	}

	if ok, _ := NewRef(Schema{Type: TypeInteger}, "AWS::StackName").Validate(template, []string{}); ok {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if ok, _ := NewRef(Schema{}, "invalid").Validate(template, []string{}); ok {
		t.Error("Should fail on invalid ref")
	}
}

func TestRefInferType(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: TypeInteger,
					},
				},
			},
		},

		Parameters: map[string]Parameter{
			"MyParameter": Parameter{
				Type: TypeBool,
			},
		},
	}

	if (Ref{target: "MyResource"}).InferType(template) != TypeInteger {
		t.Error("Ref should infer type of resource")
	}

	if (Ref{target: "MyParameter"}).InferType(template) != TypeBool {
		t.Error("Ref should infer type of parameter")
	}

	if (Ref{target: "invalid"}).InferType(template) != TypeUnknown {
		t.Error("Ref should return unknown for bad ref")
	}
}
