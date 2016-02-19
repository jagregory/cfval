package route_53

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestGeoLocationSubdivisionCodeValidation(t *testing.T) {
	template := &parse.Template{}
	res := schema.Resource{
		Properties: schema.Properties{
			"GeoLocation": schema.Schema{
				Type: geoLocation,
			},
		},
	}
	ctx := schema.NewInitialContext(template, schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	}))

	badCountryCtx := schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
		parse.NewTemplateResource("TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "AK",
				"CountryCode":     "AU",
			},
		}),
		res,
	})

	badSubdivisionCtx := schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
		parse.NewTemplateResource("TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "NSW",
				"CountryCode":     "US",
			},
		}),
		res,
	})

	goodCombinationCtx := schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
		parse.NewTemplateResource("TestResource", map[string]interface{}{
			"GeoLocation": map[string]interface{}{
				"SubdivisionCode": "AK",
				"CountryCode":     "US",
			},
		}),
		res,
	})

	if _, errs := res.Validate(goodCombinationCtx); errs != nil {
		t.Error("Period should pass on a valid state with US as the country", errs)
	}

	if _, errs := res.Validate(badSubdivisionCtx); errs == nil {
		t.Error("Period should fail on an invalid subdivision with US as the country")
	}

	if _, errs := res.Validate(badCountryCtx); errs == nil {
		t.Error("Period should fail when subdivision set without US as the country")
	}
}
