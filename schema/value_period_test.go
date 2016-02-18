package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestPeriodValidation(t *testing.T) {
	template := &parse.Template{}
	self := ResourceWithDefinition{
		parse.NewTemplateResource(template, "", make(map[string]interface{})),
		Resource{},
	}
	ctx := Context{
		Definitions: NewResourceDefinitions(nil),
		Path:        []string{},
		Template:    template,
	}

	if _, errs := Period.Validate(Schema{}, "", self, ctx); errs == nil {
		t.Error("Period should fail on empty string")
	}

	if _, errs := Period.Validate(Schema{}, "abc", self, ctx); errs == nil {
		t.Error("Period should fail on anything which isn't a period")
	}

	for _, ex := range []string{"0", "10", "119", "260"} {
		if _, errs := Period.Validate(Schema{}, ex, self, ctx); errs == nil {
			t.Errorf("Period should fail on number which isn't a multiple of 60 (ex: %s)", ex)
		}
	}

	for _, ex := range []string{"60", "120", "240"} {
		if _, errs := Period.Validate(Schema{}, ex, self, ctx); errs != nil {
			t.Errorf("Cidr should pass with numbers which are multiples of 60 (ex: %s)", ex)
		}
	}
}
