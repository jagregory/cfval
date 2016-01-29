package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestCidrValidation(t *testing.T) {
	template := schema.Template{}
	context := []string{}

	if ok, _ := cidr("", template, context); ok {
		t.Error("Cidr should fail on empty string")
	}

	if ok, _ := cidr("abc", template, context); ok {
		t.Error("Cidr should fail on anything which isn't a cidr")
	}

	if ok, _ := cidr("0.0.0.0/100", template, context); ok {
		t.Error("Cidr should fail on an invalid mask")
	}

	if ok, _ := cidr("10.200.300.10/24", template, context); ok {
		t.Error("Cidr should fail on an invalid IP")
	}

	if ok, _ := cidr("10.200.30.10/24", template, context); !ok {
		t.Error("Cidr should pass with a valid cidr")
	}
}
