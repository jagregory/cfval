package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestIPAddressValidation(t *testing.T) {
	ctx := NewContextShorthand(&parse.Template{}, NewResourceDefinitions(nil), emptyCurrentResource{}, Schema{})

	valids := []string{"1.2.3.4", "192.168.20.1", "255.255.255.255"}
	invalids := []string{"1.2.3", "1.2.3.4.5", "1923.168.20.1", "256.265.355.255"}

	for _, valid := range valids {
		if _, errs := IPAddress.Validate(valid, ctx); errs != nil {
			t.Errorf("expected %s to be valid %s", valid, errs)
		}
	}

	for _, invalid := range invalids {
		if _, errs := IPAddress.Validate(invalid, ctx); errs == nil {
			t.Errorf("expected %s to be invalid %s", invalid, errs)
		}
	}
}
