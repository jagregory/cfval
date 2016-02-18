package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestRefValidate(t *testing.T) {
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

	ctx := Context{
		Definitions: NewResourceDefinitions(map[string]Resource{
			"TestResource": Resource{
				ReturnValue: Schema{
					Type: ValueString,
				},
			},
		}),
		Path:     []string{},
		Template: template,
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource1").Validate(ctx); errs == nil {
		t.Error("Should fail on valid resource ref with Unknown ref type")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Resource2").Validate(ctx); errs != nil {
		t.Error("Should pass on valid resource ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Resource2").Validate(ctx); errs == nil {
		t.Error("Should fail on valid resource ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter1").Validate(ctx); errs == nil {
		t.Error("Should fail on valid parameter ref with Unknown ref type", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "Parameter2").Validate(ctx); errs != nil {
		t.Error("Should pass on valid parameter ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "Parameter2").Validate(ctx); errs == nil {
		t.Error("Should fail on valid parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "AWS::StackName").Validate(ctx); errs != nil {
		t.Error("Should pass on valid pseudo-parameter ref with matching types", errs)
	}

	if _, errs := NewRef(Schema{Type: ValueNumber}, "AWS::StackName").Validate(ctx); errs == nil {
		t.Error("Should fail on valid pseudo-parameter ref with non-matching types")
	}

	if _, errs := NewRef(Schema{Type: ValueString}, "invalid").Validate(ctx); errs == nil {
		t.Error("Should fail on invalid ref")
	}
}
