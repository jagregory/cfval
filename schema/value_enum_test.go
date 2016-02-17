package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestEnumValidation(t *testing.T) {
	template := &parse.Template{}
	tr := parse.NewTemplateResource(template)
	context := []string{}

	enum := EnumValue{
		Options: []string{"a", "b", "c"},
	}

	if _, errs := enum.Validate(Schema{}, "", tr, nil, context); errs == nil {
		t.Error("Enum should fail on empty string")
	}

	if _, errs := enum.Validate(Schema{}, "d", tr, nil, context); errs == nil {
		t.Error("Enum should fail on anything which isn't a valid option")
	}

	if _, errs := enum.Validate(Schema{}, "b", tr, nil, context); errs != nil {
		t.Error("Enum should pass on a valid option")
	}
}
