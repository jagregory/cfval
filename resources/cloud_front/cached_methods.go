package cloud_front

import . "github.com/jagregory/cfval/schema"

var cachedMethods = FuncType{
	Description: "CloudFront Cached Methods",

	Fn: FixedArrayValidate(
		[]string{"HEAD", "GET"},
		[]string{"GET", "HEAD", "OPTIONS"},
	),
}
