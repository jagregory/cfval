package elastic_load_balancing

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-LBCookieStickinessPolicy.html
var lbCookieStickinessPolicy = NestedResource{
	Description: "ElasticLoadBalancing LBCookieStickinessPolicy",
	Properties: Properties{
		"CookieExpirationPeriod": Schema{
			Type: ValueString,
		},

		"PolicyName": Schema{
			Type: ValueString,
		},
	},
}
