package rds

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
var securityGroupRule = NestedResource{
	Description: "RDS Security Group Rule",

	Properties: Properties{
		"CIDRIP": Schema{
			Type: CIDR,
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
