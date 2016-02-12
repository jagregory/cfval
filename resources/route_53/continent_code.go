package route_53

import . "github.com/jagregory/cfval/schema"

var continentCode = EnumValue{
	Description: "GeoLocation Continent Code",

	Options: []string{"AF", "AN", "AS", "EU", "OC", "NA", "SA"},
}
