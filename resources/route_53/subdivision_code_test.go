package route_53

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestGeoLocationSubdivisionCodeValidation(t *testing.T) {
	template := &parse.Template{}
	context := []string{}

	res := schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}

	definitions := schema.NewResourceDefinitions(map[string]func() schema.Resource{
		"TestResource": func() schema.Resource {
			return res
		},
	})

	badCountry := parse.NewTemplateResource(template)
	badCountry.Type = "TestResource"
	badCountry.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "AU",
		},
	}

	badSubdivision := parse.NewTemplateResource(template)
	badSubdivision.Type = "TestResource"
	badSubdivision.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "NSW",
			"CountryCode":     "US",
		},
	}

	goodCombination := parse.NewTemplateResource(template)
	goodCombination.Type = "TestResource"
	goodCombination.Properties = map[string]interface{}{
		"GeoLocation": map[string]interface{}{
			"SubdivisionCode": "AK",
			"CountryCode":     "US",
		},
	}

	if _, errs := res.Validate(goodCombination, definitions, context); errs != nil {
		t.Error("Period should pass on a valid state with US as the country", errs)
	}

	if _, errs := res.Validate(badSubdivision, definitions, context); errs == nil {
		t.Error("Period should fail on an invalid subdivision with US as the country")
	}

	if _, errs := res.Validate(badCountry, definitions, context); errs == nil {
		t.Error("Period should fail when subdivision set without US as the country")
	}
}
