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
	ctx := NewContextShorthand(template, NewResourceDefinitions(nil), self, Schema{})

	enum := EnumValue{
		Options: []string{"a", "b", "c"},
	}

	if _, errs := enum.Validate("", ctx); errs == nil {
		t.Error("Enum should fail on empty string")
	}

	if _, errs := enum.Validate("d", ctx); errs == nil {
		t.Error("Enum should fail on anything which isn't a valid option")
	}

	if _, errs := enum.Validate("b", ctx); errs != nil {
		t.Error("Enum should pass on a valid option")
	}
}
