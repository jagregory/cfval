package route_53

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-recordset-geolocation.html
var geoLocation = NestedResource{
	Description: "Route 53 Record Set GeoLocation",
	Properties: Properties{
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
			ValidateFunc: func(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
				if countryCode, found := ctx.CurrentResource().PropertyValueOrDefault("CountryCode"); found && countryCode != "US" {
					return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Can only be set when CountryCode is US")}
				}

				return reporting.ValidateOK, nil
			},
		},
	},
}
