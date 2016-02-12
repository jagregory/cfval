package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectionsettings.html
var connectionSettings = NestedResource{
	Description: "Elastic Load Balancing ConnectionSettings",
	Properties: Properties{
		"IdleTimeout": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
