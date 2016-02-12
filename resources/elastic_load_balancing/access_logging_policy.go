package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-accessloggingpolicy.html
var accessLoggingPolicy = NestedResource{
	Description: "Elastic Load Balancing AccessLoggingPolicy",
	Properties: Properties{
		"EmitInterval": Schema{
			Type: ValueNumber,
		},

		"Enabled": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"S3BucketName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"S3BucketPrefix": Schema{
			Type: ValueString,
		},
	},
}
