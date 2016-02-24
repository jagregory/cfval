package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html
var DBSecurityGroup = Resource{
	AwsType: "AWS::RDS::DBSecurityGroup",

	// Name
	ReturnValue: Schema{
		Type: dbSecurityGroupName,
	},

	Properties: Properties{
		"EC2VpcId": Schema{
			Type: VpcID,
		},

		"DBSecurityGroupIngress": Schema{
			Type:     Multiple(securityGroupRule),
			Required: constraints.Always,
		},

		"GroupDescription": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Tags": Schema{
			Type: Multiple(common.ResourceTag),
		},
	},
}
