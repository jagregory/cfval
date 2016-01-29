package schema

import (
	"testing"
)

func TestResourcePropertyConflictValidation(t *testing.T) {
	template := Template{}
	tr := TemplateResource{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:      TypeString,
				Conflicts: []string{"Option2"},
			},

			"Option2": Schema{
				Type:      TypeString,
				Conflicts: []string{"Option1"},
			},
		},
	}

	nothingSet := map[string]interface{}{}
	option1Set := map[string]interface{}{
		"Option1": "value",
	}
	option2Set := map[string]interface{}{
		"Option2": "value",
	}
	bothSet := map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}

	if ok, _ := resource.Validate(template, tr, nothingSet, context); !ok {
		t.Error("Resource should pass if both neither Option1 or Option2 are set")
	}

	if ok, _ := resource.Validate(template, tr, option1Set, context); !ok {
		t.Error("Resource should pass if only Option1 set")
	}

	if ok, _ := resource.Validate(template, tr, option2Set, context); !ok {
		t.Error("Resource should pass if only Option2 set")
	}

	if ok, _ := resource.Validate(template, tr, bothSet, context); ok {
		t.Error("Resource should fail if both Option1 or Option2 are set")
	}
}

func TestResourcePropertyRequiredIfValidation(t *testing.T) {
	template := Template{}
	tr := TemplateResource{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:       TypeString,
				RequiredIf: []string{"Option2"},
			},

			"Option2": Schema{
				Type: TypeString,
			},
		},
	}

	nothingSet := map[string]interface{}{}
	option1Set := map[string]interface{}{
		"Option1": "value",
	}
	option2Set := map[string]interface{}{
		"Option2": "value",
	}
	bothSet := map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}

	if ok, _ := resource.Validate(template, tr, nothingSet, context); !ok {
		t.Error("Resource should pass if neither Option1 or Option2 are set")
	}

	if ok, _ := resource.Validate(template, tr, option1Set, context); !ok {
		t.Error("Resource should pass if only Option1 set")
	}

	if ok, _ := resource.Validate(template, tr, option2Set, context); ok {
		t.Error("Resource should fail if only Option2 set")
	}

	if ok, _ := resource.Validate(template, tr, bothSet, context); !ok {
		t.Error("Resource should pass if both Option1 and Option2 are set")
	}
}

func TestResourcePropertyRequiredUnlessValidation(t *testing.T) {
	template := Template{}
	tr := TemplateResource{}
	context := []string{}

	resource := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:           TypeString,
				RequiredUnless: []string{"Option2"},
			},

			"Option2": Schema{
				Type: TypeString,
			},
		},
	}

	nothingSet := map[string]interface{}{}
	option1Set := map[string]interface{}{
		"Option1": "value",
	}
	option2Set := map[string]interface{}{
		"Option2": "value",
	}
	bothSet := map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}

	if ok, _ := resource.Validate(template, tr, nothingSet, context); ok {
		t.Error("Resource should fail if neither Option1 or Option2 are set")
	}

	if ok, _ := resource.Validate(template, tr, option1Set, context); !ok {
		t.Error("Resource should pass if only Option1 set")
	}

	if ok, _ := resource.Validate(template, tr, option2Set, context); !ok {
		t.Error("Resource should pass if only Option2 set")
	}

	if ok, _ := resource.Validate(template, tr, bothSet, context); !ok {
		t.Error("Resource should pass if both Option1 and Option2 are set")
	}
}
