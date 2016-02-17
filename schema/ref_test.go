package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestRefValidate(t *testing.T) {
	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	})

	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"Resource1": &parse.TemplateResource{},
			"Resource2": &parse.TemplateResource{
				Type: "TestResource",
			},
		},
		Parameters: map[string]parse.Parameter{
			"Parameter1": parse.Parameter{},
			"Parameter2": parse.Parameter{
				Type: "String",
			},
		},
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource1").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on valid resource ref with Unknown ref type")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource2").Validate(template, definitions, []string{}); errs != nil {
		t.Error("Should pass on valid resource ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Resource2").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter1").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on valid parameter ref with Unknown ref type", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter2").Validate(template, definitions, []string{}); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Parameter2").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "AWS::StackName").Validate(template, definitions, []string{}); errs != nil {
		t.Error("Should pass on valid pseudo-parameter ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "AWS::StackName").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "invalid").Validate(template, definitions, []string{}); errs == nil {
		t.Error("Should fail on invalid ref")
	}
}
