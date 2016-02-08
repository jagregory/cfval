package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

// SubdivisionCode is the first to use HasProperty for looking up if a specific
// value is set, so we should test it actually works.
func TestGeoLocationSubdivisionCodeValidation(t *testing.T) {
	template := &schema.Template{}
	context := []string{}

	badCountry := schema.NewTemplateResource(template)
	badCountry.Definition = schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}
	badCountry.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "AU",
		},
	}

	badSubdivision := schema.NewTemplateResource(template)
	badSubdivision.Definition = schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}
	badSubdivision.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "NSW",
			"CountryCode":     "US",
		},
	}

	goodCombination := schema.NewTemplateResource(template)
	goodCombination.Definition = schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}
	goodCombination.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "US",
		},
	}

	if _, errs := goodCombination.Validate(context); errs != nil {
		t.Error("Period should pass on a valid state with US as the country", errs)
	}

	if _, errs := badSubdivision.Validate(context); errs == nil {
		t.Error("Period should fail on an invalid subdivision with US as the country")
	}

	if _, errs := badCountry.Validate(context); errs == nil {
		t.Error("Period should fail when subdivision set without US as the country")
	}
}
