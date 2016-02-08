package schema

import (
	"encoding/json"
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

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
		Type: TypeAvailabilityZoneName,
	},

	"AWS::EC2::Image::Id": Schema{
		Type: TypeImageId,
	},

	"AWS::EC2::Instance::Id": Schema{
		Type: TypeInstanceId,
	},

	"AWS::EC2::KeyPair::KeyName": Schema{
		Type: TypeKeyName,
	},

	"AWS::EC2::SecurityGroup::GroupName": Schema{
		Type: TypeSecurityGroupName,
	},

	"AWS::EC2::SecurityGroup::Id": Schema{
		Type: TypeSecurityGroupId,
	},

	"AWS::EC2::Subnet::Id": Schema{
		Type: TypeSubnetId,
	},

	"AWS::EC2::Volume::Id": Schema{
		Type: TypeVolumeId,
	},

	"AWS::EC2::VPC::Id": Schema{
		Type: TypeVPCId,
	},

	"AWS::Route53::HostedZone::Id": Schema{
		Type: TypeHostedZoneId,
	},

	"List<AWS::EC2::AvailabilityZone::Name>": Schema{
		Type:  TypeAvailabilityZoneName,
		Array: true,
	},

	"List<AWS::EC2::Image::Id>": Schema{
		Type:  TypeImageId,
		Array: true,
	},

	"List<AWS::EC2::Instance::Id>": Schema{
		Type:  TypeInstanceId,
		Array: true,
	},

	"List<AWS::EC2::SecurityGroup::GroupName>": Schema{
		Type:  TypeSecurityGroupName,
		Array: true,
	},

	"List<AWS::EC2::SecurityGroup::Id>": Schema{
		Type:  TypeSecurityGroupId,
		Array: true,
	},

	"List<AWS::EC2::Subnet::Id>": Schema{
		Type:  TypeSubnetId,
		Array: true,
	},

	"List<AWS::EC2::Volume::Id>": Schema{
		Type:  TypeVolumeId,
		Array: true,
	},

	"List<AWS::EC2::VPC::Id>": Schema{
		Type:  TypeVPCId,
		Array: true,
	},

	"List<AWS::Route53::HostedZone::Id>": Schema{
		Type:  TypeHostedZoneId,
		Array: true,
	},
}

type Parameter struct {
	Schema
}

func (p *Parameter) UnmarshalJSON(b []byte) (err error) {
	temp := struct {
		Type string
	}{}

	if err = json.Unmarshal(b, &temp); err != nil {
		panic("Unexpected type unmarshalling parameter")
	}

	if s, found := parameterTypeSchemas[temp.Type]; found {
		p.Schema = s
		return nil
	}

	return fmt.Errorf("Unexpected type for Parameter %s", temp.Type)
}

func (Parameter) Validate([]string) (bool, []reporting.Failure) {
	return true, nil
}

func (p Parameter) TargetType() ValueType {
	if t, ok := p.Type.(ValueType); ok {
		return t
	}

	return ValueUnknown
}
