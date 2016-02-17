package rds

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-security-group-ingress.html
var DBSecurityGroupIngress = Resource{
	AwsType: "AWS::RDS::DBSecurityGroupIngress",

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CIDRIP": Schema{
			Type: CIDR,
		},

		"DBSecurityGroupName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"EC2SecurityGroupId": Schema{
			Type: SecurityGroupID,
		},

		"EC2SecurityGroupName": Schema{
			Type: SecurityGroupName,
		},

		"EC2SecurityGroupOwnerId": Schema{
			Type: ValueString,
		},
	},
}
