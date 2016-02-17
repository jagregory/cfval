package schema

import (
	"fmt"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type ResourceDefinitions interface {
	Lookup(awsType string) Resource
	LookupParameter(awsType string) Schema
}

func NewResourceDefinitions(definitions map[string]func() Resource) ResourceDefinitions {
	if definitions == nil {
		definitions = make(map[string]func() Resource)
	}

	return externalResourceDefinitions{
		definitions: definitions,
	}
}

type externalResourceDefinitions struct {
	definitions map[string]func() Resource
}

func (e externalResourceDefinitions) Lookup(awsType string) Resource {
	ctor := e.definitions[awsType]

	if ctor == nil {
		return Resource{
			ValidateFunc: func(tr parse.TemplateResource, context []string) (reporting.ValidateResult, reporting.Reports) {
				return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), context)}
			},
		}
	}

	return ctor()
}

func (externalResourceDefinitions) LookupParameter(awsType string) Schema {
	return parameterTypeSchemas[awsType]
}

var parameterTypeSchemas = map[string]Schema{
	"String": Schema{
		Type: ValueString,
	},

	"Number": Schema{
		Type: ValueNumber,
	},

	"List<Number>": Schema{
		Type:  ValueNumber,
		Array: true,
	},

	"CommaDelimitedList": Schema{
		Type:  ValueString,
		Array: true,
	},

	"AWS::EC2::AvailabilityZone::Name": Schema{
		Type: AvailabilityZone,
	},

	"AWS::EC2::Image::Id": Schema{
		Type: ImageID,
	},

	"AWS::EC2::Instance::Id": Schema{
		Type: InstanceID,
	},

	"AWS::EC2::KeyPair::KeyName": Schema{
		Type: KeyName,
	},

	"AWS::EC2::SecurityGroup::GroupName": Schema{
		Type: SecurityGroupName,
	},

	"AWS::EC2::SecurityGroup::Id": Schema{
		Type: SecurityGroupID,
	},

	"AWS::EC2::Subnet::Id": Schema{
		Type: SubnetID,
	},

	"AWS::EC2::Volume::Id": Schema{
		Type: VolumeID,
	},

	"AWS::EC2::VPC::Id": Schema{
		Type: VpcID,
	},

	"AWS::Route53::HostedZone::Id": Schema{
		Type: HostedZoneID,
	},

	"List<AWS::EC2::AvailabilityZone::Name>": Schema{
		Type:  AvailabilityZone,
		Array: true,
	},

	"List<AWS::EC2::Image::Id>": Schema{
		Type:  ImageID,
		Array: true,
	},

	"List<AWS::EC2::Instance::Id>": Schema{
		Type:  InstanceID,
		Array: true,
	},

	"List<AWS::EC2::SecurityGroup::GroupName>": Schema{
		Type:  SecurityGroupName,
		Array: true,
	},

	"List<AWS::EC2::SecurityGroup::Id>": Schema{
		Type:  SecurityGroupID,
		Array: true,
	},

	"List<AWS::EC2::Subnet::Id>": Schema{
		Type:  SubnetID,
		Array: true,
	},

	"List<AWS::EC2::Volume::Id>": Schema{
		Type:  VolumeID,
		Array: true,
	},

	"List<AWS::EC2::VPC::Id>": Schema{
		Type:  VpcID,
		Array: true,
	},

	"List<AWS::Route53::HostedZone::Id>": Schema{
		Type:  HostedZoneID,
		Array: true,
	},
}
