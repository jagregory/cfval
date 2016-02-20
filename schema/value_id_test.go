package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestIDValidation(t *testing.T) {
	ctx := NewContextShorthand(&parse.Template{}, NewResourceDefinitions(nil), emptyCurrentResource{}, Schema{})
	id := newResourceID("TestID", "test", false)

	valids := []string{"test-15a4417c", "test-12345678", "test-abcdefgh"}
	invalids := []string{"rest-15a4417c", "test-15a4417c1234", "test-15a4417c1234abcf8", "test-abc"}

	for _, valid := range valids {
		if _, errs := id.Validate(valid, ctx); errs != nil {
			t.Errorf("expected %s to be valid %s", valid, errs)
		}
	}

	for _, invalid := range invalids {
		if _, errs := id.Validate(invalid, ctx); errs == nil {
			t.Errorf("expected %s to be invalid %s", invalid, errs)
		}
	}
}

func TestLongIDValidation(t *testing.T) {
	ctx := NewContextShorthand(&parse.Template{}, NewResourceDefinitions(nil), emptyCurrentResource{}, Schema{})
	id := newResourceID("TestID", "test", true)

	valids := []string{"test-15a4417c", "test-12345678", "test-abcdefgh", "test-15a4417c1234abcf8"}
	invalids := []string{"rest-15a4417c", "rest-15a4417c1234abcf8", "test-15a4417c1234ab", "test-abc"}

	for _, valid := range valids {
		if _, errs := id.Validate(valid, ctx); errs != nil {
			t.Errorf("expected %s to be valid %s", valid, errs)
		}
	}

	for _, invalid := range invalids {
		if _, errs := id.Validate(invalid, ctx); errs == nil {
			t.Errorf("expected %s to be invalid %s", invalid, errs)
		}
	}
}
