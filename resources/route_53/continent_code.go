package route_53

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets_Requests.html#change-rrsets-request-continent-code
var continentCode = EnumValue{
	Description: "GeoLocation Continent Code",

	Options: []string{"AF", "AN", "AS", "EU", "OC", "NA", "SA"},
}
