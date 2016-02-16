package rds

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group.html
func DBSecurityGroup() Resource {
	return Resource{
		AwsType: "AWS::RDS::DBSecurityGroup",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"EC2VpcId": Schema{
				Type: VpcID,
			},

			"DBSecurityGroupIngress": Schema{
				Type:     securityGroupRule,
				Array:    true,
				Required: constraints.Always,
			},

			"GroupDescription": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Tags": Schema{
				Type:  common.ResourceTag,
				Array: true,
			},
		},
	}
}
