package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

// SubdivisionCode is the first to use HasProperty for looking up if a specific
// value is set, so we should test it actually works.
func TestGeoLocationSubdivisionCodeValidation(t *testing.T) {
	template := schema.Template{}
	context := []string{}

	badCountry := schema.TemplateResource{
		Definition: geoLocation,
		Properties: map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "AU",
		},
	}
	badSubdivision := schema.TemplateResource{
		Definition: geoLocation,
		Properties: map[string]interface{}{
			"SubdivisionCode": "NSW",
			"CountryCode":     "US",
		},
	}
	goodCombination := schema.TemplateResource{
		Definition: geoLocation,
		Properties: map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "US",
		},
	}

	if ok, _ := geoLocation.Validate(template, goodCombination, goodCombination.Properties, context); !ok {
		t.Error("Period should pass on a valid state with US as the country")
	}

	if ok, _ := geoLocation.Validate(template, badSubdivision, badSubdivision.Properties, context); ok {
		t.Error("Period should fail on an invalid subdivision with US as the country")
	}

	if ok, _ := geoLocation.Validate(template, badCountry, badCountry.Properties, context); ok {
		t.Error("Period should fail when subdivision set without US as the country")
	}
}
