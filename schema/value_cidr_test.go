package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestCidrValidation(t *testing.T) {
	template := &parse.Template{}
	self := ResourceWithDefinition{
		parse.NewTemplateResource(template, "", make(map[string]interface{})),
		Resource{},
	}
	ctx := Context{
		Template: template,
		Path:     []string{},
	}

	if _, errs := CIDR.Validate(Schema{}, "", self, nil, ctx); errs == nil {
		t.Error("Cidr should fail on empty str, niling")
	}

	if _, errs := CIDR.Validate(Schema{}, "abc", self, nil, ctx); errs == nil {
		t.Error("Cidr should fail on anything which isn't a cidr")
	}

	if _, errs := CIDR.Validate(Schema{}, "0.0.0.0/100", self, nil, ctx); errs == nil {
		t.Error("Cidr should fail on an invalid mask")
	}

	if _, errs := CIDR.Validate(Schema{}, "10.200.300.10/24", self, nil, ctx); errs == nil {
		t.Error("Cidr should fail on an invalid IP")
	}

	if _, errs := CIDR.Validate(Schema{}, "10.200.30.10/24", self, nil, ctx); errs != nil {
		t.Error("Cidr should pass with a valid cidr")
	}
}
