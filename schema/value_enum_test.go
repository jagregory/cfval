package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestEnumValidation(t *testing.T) {
	template := &parse.Template{}
	self := ResourceWithDefinition{
		parse.NewTemplateResource(template, "", make(map[string]interface{})),
		Resource{},
	}
	ctx := Context{
		Template: template,
		Path:     []string{},
	}

	enum := EnumValue{
		Options: []string{"a", "b", "c"},
	}

	if _, errs := enum.Validate(Schema{}, "", self, nil, ctx); errs == nil {
		t.Error("Enum should fail on empty string")
	}

	if _, errs := enum.Validate(Schema{}, "d", self, nil, ctx); errs == nil {
		t.Error("Enum should fail on anything which isn't a valid option")
	}

	if _, errs := enum.Validate(Schema{}, "b", self, nil, ctx); errs != nil {
		t.Error("Enum should pass on a valid option")
	}
}
