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
	ctx := NewContextShorthand(template, NewResourceDefinitions(nil), self, Schema{})

	if _, errs := CIDR.Validate("", ctx); errs == nil {
		t.Error("Cidr should fail on empty str, niling")
	}

	if _, errs := CIDR.Validate("abc", ctx); errs == nil {
		t.Error("Cidr should fail on anything which isn't a cidr")
	}

	if _, errs := CIDR.Validate("0.0.0.0/100", ctx); errs == nil {
		t.Error("Cidr should fail on an invalid mask")
	}

	if _, errs := CIDR.Validate("10.200.300.10/24", ctx); errs == nil {
		t.Error("Cidr should fail on an invalid IP")
	}

	if _, errs := CIDR.Validate("10.200.30.10/24", ctx); errs != nil {
		t.Error("Cidr should pass with a valid cidr")
	}
}
