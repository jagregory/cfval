package route_53

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestGeoLocationSubdivisionCodeValidation(t *testing.T) {
	template := &parse.Template{}
	ctx := schema.Context{
		Template: template,
		Path:     []string{},
	}

	res := schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}

	definitions := schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	})

	badCountry := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "AK",
				"CountryCode":     "AU",
			},
		}),
		res,
	}

	badSubdivision := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "NSW",
				"CountryCode":     "US",
			},
		}),
		res,
	}

	goodCombination := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "AK",
				"CountryCode":     "US",
			},
		}),
		res,
	}

	if _, errs := res.Validate(goodCombination, definitions, ctx); errs != nil {
		t.Error("Period should pass on a valid state with US as the country", errs)
	}

	if _, errs := res.Validate(badSubdivision, definitions, ctx); errs == nil {
		t.Error("Period should fail on an invalid subdivision with US as the country")
	}

	if _, errs := res.Validate(badCountry, definitions, ctx); errs == nil {
		t.Error("Period should fail when subdivision set without US as the country")
	}
}
