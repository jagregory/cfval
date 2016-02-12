package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html
var appCookieStickinessPolicy = NestedResource{
	Description: "ElasticLoadBalancing AppCookieStickinessPolicy",
	Properties: Properties{
		"CookieName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PolicyName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
