package schema

import (
	"fmt"
	"regexp"

	"github.com/jagregory/cfval/reporting"
)

var AllocationID = newResourceID("EipAllocationID", "eipalloc", ShortID)
var ImageID = newResourceID("AMI", "ami", ShortID)
var InstanceID = newResourceID("InstanceID", "i", LongID)
var InternetGatewayID = newResourceID("InternetGatewayID", "igw", ShortID)
var NetworkAclID = newResourceID("NetworkAclID", "acl", ShortID)
var NetworkInterfaceID = newResourceID("NetworkInterfaceID", "eni", ShortID)
var RouteTableID = newResourceID("RouteTableID", "rtb", ShortID)
var SecurityGroupID = newResourceID("SecurityGroupID", "sg", ShortID)
var SnapshotID = newResourceID("SnapshotID", "vol", LongID)
var SubnetID = newResourceID("SubnetID", "subnet", ShortID)
var VolumeID = newResourceID("VolumeID", "vol", LongID)
var VpcID = newResourceID("VpcID", "vpc", ShortID)
var VpcPeeringConnectionID = newResourceID("VpcPeeringConnectionID", "pcx", ShortID)
var VpnConnectionID = newResourceID("VpnConnectionID", "vpn", ShortID)
var VpnGatewayID = newResourceID("VpnGatewayID", "vgw", ShortID)

const (
	LongID  bool = true
	ShortID      = false
)

type resourceID struct {
	description, example string
	regex                *regexp.Regexp
}

func (id resourceID) Validate(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if result, errs := ValueString.Validate(value, ctx); result == reporting.ValidateAbort || errs != nil {
		return reporting.ValidateOK, errs
	}

	if !id.regex.MatchString(value.(string)) {
		return reporting.ValidateOK, reporting.Reports{
			reporting.NewFailure(ctx, "Value '%s' is not a valid %s, format: %s", value, id.description, id.example),
		}
	}

	return reporting.ValidateOK, nil
}

func (resourceID) IsArray() bool {
	return false
}

func (id resourceID) Describe() string {
	return id.description
}

func (resourceID) PropertyDefault(string) (interface{}, bool) {
	return nil, false
}

func (id resourceID) Same(to PropertyType) bool {
	if toID, ok := to.(resourceID); ok {
		return toID.description == id.description
	}

	return false
}

func (id resourceID) CoercibleTo(to PropertyType) Coercion {
	if to == ValueString {
		return CoercionAlways
	} else if rid, ok := to.(resourceID); ok && rid.Describe() == id.Describe() {
		return CoercionAlways
	} else if to == ValueUnknown {
		return CoercionBegrudgingly
	}

	return CoercionNever
}

func newResourceID(description, prefix string, long bool) resourceID {
	var regex *regexp.Regexp
	var example string

	if long {
		regex = regexp.MustCompile(fmt.Sprintf("^%s-([a-z0-9]{8}|[a-z0-9]{17})$", prefix))
		example = fmt.Sprintf("%s-5fe31a21 (\"%s-\" followed by 8 alphanumeric characters) or %s-5fe31a21663d3f1c4 (\"%s-\" followed by 17 alphanumeric characters)", prefix, prefix)
	} else {
		regex = regexp.MustCompile(fmt.Sprintf("^%s-[a-z0-9]{8}$", prefix))
		example = fmt.Sprintf("%s-5fe31a21 (\"%s-\" followed by 8 alphanumeric characters)", prefix, prefix)
	}

	return resourceID{
		description: description,
		example:     example,
		regex:       regex,
	}
}
