package route_53

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

var geoLocation = NestedResource{
	Description: "Route 53 Record Set GeoLocation",
	Properties: map[string]Schema{
		"ContinentCode": Schema{
			Type:     continentCode,
			Required: constraints.PropertyNotExists("CountryCode"),
			Conflicts: constraints.Any{
				constraints.PropertyExists("CountryCode"),
				constraints.PropertyExists("SubdivisionCode"),
			},
		},

		"CountryCode": Schema{
			Type:      common.CountryCode,
			Required:  constraints.PropertyNotExists("ContinentCode"),
			Conflicts: constraints.PropertyExists("ContinentCode"),
		},

		"SubdivisionCode": Schema{
			Type:      subdivisionCode,
			Conflicts: constraints.PropertyExists("ContinentCode"),
			ValidateFunc: func(property Schema, value interface{}, self SelfRepresentation, context []string) (reporting.ValidateResult, reporting.Failures) {
				if countryCode, found := self.Property("CountryCode"); found && countryCode != "US" {
					return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Can only be set when CountryCode is US", context)}
				}

				return reporting.ValidateOK, nil
			},
		},
	},
}
