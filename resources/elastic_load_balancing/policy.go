package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html
var policy = NestedResource{
	Description: "ElasticLoadBalancing Policy",
	Properties: Properties{
		"Attributes": Schema{
			Type:     common.NameValue,
			Required: constraints.Always,
			Array:    true,
		},

		"InstancePorts": Schema{
			Type:  ValueString,
			Array: true,
		},

		"LoadBalancerPorts": Schema{
			Type:  ValueString,
			Array: true,
		},

		"PolicyName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PolicyType": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}
