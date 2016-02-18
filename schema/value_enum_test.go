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
	context := []string{}

	enum := EnumValue{
		Options: []string{"a", "b", "c"},
	}

	if _, errs := enum.Validate(Schema{}, "", self, template, nil, context); errs == nil {
		t.Error("Enum should fail on empty string")
	}

	if _, errs := enum.Validate(Schema{}, "d", self, template, nil, context); errs == nil {
		t.Error("Enum should fail on anything which isn't a valid option")
	}

	if _, errs := enum.Validate(Schema{}, "b", self, template, nil, context); errs != nil {
		t.Error("Enum should pass on a valid option")
	}
}
