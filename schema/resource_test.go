package schema

import (
	"testing"
)

func TestResourcePropertyConflictValidation(t *testing.T) {
	template := &Template{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:      ValueString,
				Conflicts: PropertyExists("Option2"),
			},

			"Option2": Schema{
				Type:      ValueString,
				Conflicts: PropertyExists("Option1"),
			},
		},
	}

	nothingSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{}}
	option1Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := nothingSet.Validate(context); errs != nil {
		t.Error("Resource should pass if both neither Option1 or Option2 are set", errs)
	}

	if _, errs := option1Set.Validate(context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := option2Set.Validate(context); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := bothSet.Validate(context); errs == nil {
		t.Error("Resource should fail if both Option1 or Option2 are set")
	}
}

func TestResourcePropertyRequiredIfValidation(t *testing.T) {
	template := &Template{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:       ValueString,
				RequiredIf: PropertyExists("Option2"),
			},

			"Option2": Schema{
				Type: ValueString,
			},
		},
	}

	nothingSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{}}
	option1Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := nothingSet.Validate(context); errs != nil {
		t.Error("Resource should pass if neither Option1 or Option2 are set", errs)
	}

	if _, errs := option1Set.Validate(context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := option2Set.Validate(context); errs == nil {
		t.Error("Resource should fail if only Option2 set")
	}

	if _, errs := bothSet.Validate(context); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}

func TestResourcePropertyRequiredUnlessValidation(t *testing.T) {
	template := &Template{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:           ValueString,
				RequiredUnless: PropertyExists("Option2"),
			},

			"Option2": Schema{
				Type: ValueString,
			},
		},
	}

	nothingSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{}}
	option1Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := TemplateResource{template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := nothingSet.Validate(context); errs == nil {
		t.Error("Resource should fail if neither Option1 or Option2 are set")
	}

	if _, errs := option1Set.Validate(context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := option2Set.Validate(context); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := bothSet.Validate(context); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}
