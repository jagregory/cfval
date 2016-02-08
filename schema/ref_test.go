package schema

import "testing"

func TestRefValidate(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"Resource1": TemplateResource{},
			"Resource2": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: ValueString,
					},
				},
			},
		},
		Parameters: map[string]Parameter{
			"Parameter1": Parameter{},
			"Parameter2": Parameter{
				Schema: Schema{
					Type: ValueString,
				},
			},
		},
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource1").Validate(template, []string{}); errs != nil {
		t.Error("Should pass on valid resource ref with Unknown ref type")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource2").Validate(template, []string{}); errs != nil {
		t.Error("Should pass on valid resource ref with matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Resource2").Validate(template, []string{}); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter1").Validate(template, []string{}); errs != nil {
		t.Error("Should pass on valid parameter ref with Unknown ref type")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter2").Validate(template, []string{}); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Parameter2").Validate(template, []string{}); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "AWS::StackName").Validate(template, []string{}); errs != nil {
		t.Error("Should pass on valid pseudo-parameter ref with matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "AWS::StackName").Validate(template, []string{}); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{}, "invalid").Validate(template, []string{}); errs == nil {
		t.Error("Should fail on invalid ref")
	}
}

func TestRefInferType(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{
				Definition: Resource{
					ReturnValue: Schema{
						Type: ValueNumber,
					},
				},
			},
		},

		Parameters: map[string]Parameter{
			"MyParameter": Parameter{
				Schema: Schema{
					Type: ValueBool,
				},
			},
		},
	}

	if (Ref{target: "MyResource"}).InferType(template) != ValueNumber {
		t.Error("Ref should infer type of resource")
	}

	if (Ref{target: "MyParameter"}).InferType(template) != ValueBool {
		t.Error("Ref should infer type of parameter")
	}

	if (Ref{target: "invalid"}).InferType(template) != ValueUnknown {
		t.Error("Ref should return unknown for bad ref")
	}
}
