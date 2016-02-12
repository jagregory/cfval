package cloud_front

import . "github.com/jagregory/cfval/schema"

var priceClass = EnumValue{
	Description: "CloudFront PriceClass",

	Options: []string{"PriceClass_All", "PriceClass_200", "PriceClass_100"},
}
