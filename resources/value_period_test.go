package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestPeriodValidation(t *testing.T) {
	template := schema.Template{}
	tr := schema.TemplateResource{}
	context := []string{}

	if ok, _ := period("", template, tr, context); ok {
		t.Error("Period should fail on empty string")
	}

	if ok, _ := period("abc", template, tr, context); ok {
		t.Error("Period should fail on anything which isn't a period")
	}

	for _, ex := range []string{"0", "10", "119", "260"} {
		if ok, _ := period(ex, template, tr, context); ok {
			t.Errorf("Period should fail on number which isn't a multiple of 60 (ex: %s)", ex)
		}
	}

	for _, ex := range []string{"60", "120", "240"} {
		if ok, _ := period(ex, template, tr, context); !ok {
			t.Errorf("Cidr should pass with numbers which are multiples of 60 (ex: %s)", ex)
		}
	}
}
