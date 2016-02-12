package auto_scaling

import . "github.com/jagregory/cfval/schema"

var placementTenancy = EnumValue{
	Description: "Launch Configuration Placement Tenancy",

	Options: []string{"default", "dedicated"},
}
