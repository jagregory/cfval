package schema

import (
	"testing"
)

func TestOutputValidation(t *testing.T) {
	template := &Template{
		Outputs: map[string]Output{},
	}
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

	nothingSet := TemplateResource{Template: template, Definition: resource, Properties: map[string]interface{}{}}
	option1Set := TemplateResource{Template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := TemplateResource{Template: template, Definition: resource, Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := TemplateResource{Template: template, Definition: resource, Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if ok, _ := nothingSet.Validate(context); !ok {
		t.Error("Resource should pass if both neither Option1 or Option2 are set")
	}

	if ok, _ := option1Set.Validate(context); !ok {
		t.Error("Resource should pass if only Option1 set")
	}

	if ok, _ := option2Set.Validate(context); !ok {
		t.Error("Resource should pass if only Option2 set")
	}

	if ok, _ := bothSet.Validate(context); ok {
		t.Error("Resource should fail if both Option1 or Option2 are set")
	}
}
