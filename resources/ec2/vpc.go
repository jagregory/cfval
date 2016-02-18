package ec2

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html
var VPC = Resource{
	AwsType: "AWS::EC2::VPC",

	Attributes: map[string]Schema{
		"CidrBlock": Schema{
			Type: CIDR,
		},

		"DefaultNetworkAcl": Schema{
			Type: ValueString,
		},

		"DefaultSecurityGroup": Schema{
			Type: SecurityGroupID,
		},
	},

	// ID
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"CidrBlock": Schema{
			Type:     CIDR,
			Required: constraints.Always,
		},

		"EnableDnsSupport": Schema{
			Type:    ValueBool,
			Default: true,
		},

		"EnableDnsHostnames": Schema{
			Type:    ValueBool,
			Default: false,
			ValidateFunc: func(property Schema, value interface{}, self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.ValidateResult, reporting.Reports) {
				if enableDnsSupport, _ := self.PropertyValue("EnableDnsSupport"); value == true && enableDnsSupport == false {
					return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("You can only set EnableDnsHostnames to true if you also set the EnableDnsSupport attribute to true.", path)}
				}

				return reporting.ValidateOK, nil
			},
		},

		"InstanceTenancy": Schema{
			Type: EnumValue{
				Description: "InstanceTenancy",
				Options:     []string{"default", "dedicated"},
			},
		},

		"Tags": Schema{
			Type:  common.ResourceTag,
			Array: true,
		},
	},
}
