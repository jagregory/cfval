package schema

import (
	"testing"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
)

func TestResourcePropertyConflictValidation(t *testing.T) {
	template := &parse.Template{}
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
	ctx := NewInitialContext(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}))

	nothingSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{}),
		res,
	}
	option1Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
		}),
		res,
	}
	option2Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option2": "value",
		}),
		res,
	}
	bothSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
			"Option2": "value",
		}),
		res,
	}

	if _, errs := res.Validate(NewResourceContext(ctx, nothingSet)); errs != nil {
		t.Error("Resource should pass if both neither Option1 or Option2 are set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option1Set)); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option2Set)); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, bothSet)); errs == nil {
		t.Error("Resource should fail if both Option1 or Option2 are set")
	}
}

func TestSchemaRequiredValidation(t *testing.T) {
	template := &parse.Template{}
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
	ctx := NewInitialContext(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}))

	nothingSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{}),
		res,
	}
	option1Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
		}),
		res,
	}
	option2Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option2": "value",
		}),
		res,
	}
	option3Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option3": "value",
		}),
		res,
	}
	allSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
			"Option2": "value",
			"Option3": "value",
		}),
		res,
	}

	if _, errs := res.Validate(NewResourceContext(ctx, nothingSet)); errs == nil {
		t.Error("Resource should fail if Option1 isn't set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option1Set)); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option2Set)); errs == nil {
		t.Error("Resource should fail if only Option2 set")
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option3Set)); errs == nil {
		t.Error("Resource should fail if only Option3 set")
	}

	if _, errs := res.Validate(NewResourceContext(ctx, allSet)); errs != nil {
		t.Error("Resource should pass if Option1 is set with others", errs)
	}
}

func TestResourcePropertyRequiredIfValidation(t *testing.T) {
	template := &parse.Template{}
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
	ctx := NewInitialContext(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}))

	nothingSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{}),
		res,
	}
	option1Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
		}),
		res,
	}
	option2Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option2": "value",
		}),
		res,
	}
	bothSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
			"Option2": "value",
		}),
		res,
	}

	if _, errs := res.Validate(NewResourceContext(ctx, nothingSet)); errs != nil {
		t.Error("Resource should pass if neither Option1 or Option2 are set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option1Set)); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option2Set)); errs == nil {
		t.Error("Resource should fail if only Option2 set")
	}

	if _, errs := res.Validate(NewResourceContext(ctx, bothSet)); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}

func TestResourcePropertyRequiredUnlessValidation(t *testing.T) {
	template := &parse.Template{}
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
	ctx := NewInitialContext(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	}))

	nothingSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{}),
		res,
	}
	option1Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
		}),
		res,
	}
	option2Set := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option2": "value",
		}),
		res,
	}
	bothSet := ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"Option1": "value",
			"Option2": "value",
		}),
		res,
	}

	if _, errs := res.Validate(NewResourceContext(ctx, nothingSet)); errs == nil {
		t.Error("Resource should fail if neither Option1 or Option2 are set")
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option1Set)); errs != nil {
		t.Error("Resource should pass if only Option1 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, option2Set)); errs != nil {
		t.Error("Resource should pass if only Option2 set", errs)
	}

	if _, errs := res.Validate(NewResourceContext(ctx, bothSet)); errs != nil {
		t.Error("Resource should pass if both Option1 and Option2 are set", errs)
	}
}
