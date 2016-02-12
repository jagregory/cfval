package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
var connectionDrainingPolicy = NestedResource{
	Description: "Elastic Load Balancing ConnectionDrainingPolicy",
	Properties: Properties{
		"Enabled": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"Timeout": Schema{
			Type: ValueNumber,
		},
	},
}
