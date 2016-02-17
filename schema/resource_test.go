package schema

import (
	"testing"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
)

func TestResourcePropertyConflictValidation(t *testing.T) {
	template := &parse.Template{}
	context := []string{}

	res := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:      ValueString,
				Conflicts: constraints.PropertyExists("Option2"),
			},

			"Option2": Schema{
				Type:      ValueString,
				Conflicts: constraints.PropertyExists("Option1"),
			},
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	nothingSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{}}
	option1Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := res.Validate(nothingSet, definitions, context); errs != nil {
		t.Error("Resource should pass if both neither Option1 or Option2 are set", errs)
	}

	if _, errs := res.Validate(option1Set, definitions, context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(option2Set, definitions, context); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := res.Validate(bothSet, definitions, context); errs == nil {
		t.Error("Resource should fail if both Option1 or Option2 are set")
	}
}

func TestSchemaRequiredValidation(t *testing.T) {
	template := &parse.Template{}
	ctx := []string{}

	res := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Option2": Schema{
				Type:     ValueString,
				Required: constraints.Never,
			},

			"Option3": Schema{
				Type: ValueString,
			},
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	nothingSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{}}
	option1Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option2": "value",
	}}
	option3Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option3": "value",
	}}
	allSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
		"Option3": "value",
	}}

	if _, errs := res.Validate(nothingSet, definitions, ctx); errs == nil {
		t.Error("Resource should fail if Option1 isn't set", errs)
	}

	if _, errs := res.Validate(option1Set, definitions, ctx); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(option2Set, definitions, ctx); errs == nil {
		t.Error("Resource should fail if only Option2 set")
	}

	if _, errs := res.Validate(option3Set, definitions, ctx); errs == nil {
		t.Error("Resource should fail if only Option3 set")
	}

	if _, errs := res.Validate(allSet, definitions, ctx); errs != nil {
		t.Error("Resource should pass if Option1 is set with others", errs)
	}
}

func TestResourcePropertyRequiredIfValidation(t *testing.T) {
	template := &parse.Template{}
	context := []string{}

	res := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:     ValueString,
				Required: constraints.PropertyExists("Option2"),
			},

			"Option2": Schema{
				Type: ValueString,
			},
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	nothingSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{}}
	option1Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := res.Validate(nothingSet, definitions, context); errs != nil {
		t.Error("Resource should pass if neither Option1 or Option2 are set", errs)
	}

	if _, errs := res.Validate(option1Set, definitions, context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(option2Set, definitions, context); errs == nil {
		t.Error("Resource should fail if only Option2 set")
	}

	if _, errs := res.Validate(bothSet, definitions, context); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}

func TestResourcePropertyRequiredUnlessValidation(t *testing.T) {
	template := &parse.Template{}
	context := []string{}

	res := Resource{
		Properties: map[string]Schema{
			"Option1": Schema{
				Type:     ValueString,
				Required: constraints.PropertyNotExists("Option2"),
			},

			"Option2": Schema{
				Type: ValueString,
			},
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	nothingSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{}}
	option1Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
	}}
	option2Set := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option2": "value",
	}}
	bothSet := parse.TemplateResource{Tmpl: template, Type: "TestResource", Properties: map[string]interface{}{
		"Option1": "value",
		"Option2": "value",
	}}

	if _, errs := res.Validate(nothingSet, definitions, context); errs == nil {
		t.Error("Resource should fail if neither Option1 or Option2 are set")
	}

	if _, errs := res.Validate(option1Set, definitions, context); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(option2Set, definitions, context); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := res.Validate(bothSet, definitions, context); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}
