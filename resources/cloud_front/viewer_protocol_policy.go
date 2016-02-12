package cloud_front

import . "github.com/jagregory/cfval/schema"

var viewerProtocolPolicy = EnumValue{
	Description: "CloudFront ViewerProtocolPolicy",

	Options: []string{"allow-all", "redirect-to-https", "https"},
}
