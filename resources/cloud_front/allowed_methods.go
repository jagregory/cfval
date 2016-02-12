package cloud_front

import . "github.com/jagregory/cfval/schema"

var allowedMethods = FuncType{
	Description: "CloudFront Allowed Methods",

	Fn: FixedArrayValidate(
		[]string{"HEAD", "GET"},
		[]string{"GET", "HEAD", "OPTIONS"},
		[]string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
	),
}
