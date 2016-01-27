package main

type Aws_Ec2_Subnet struct {
	AvailabilityZone    AvailabilityZone `update:"replace"`
	CidrBlock           Cidr             `required:true update:"replace"`
	MapPublicIpOnLaunch BoolOrBuiltinFns
	Tags                ResourceTags
	VpcId               VpcId `required:true update:"replace"`
}

// TODO: I think these could be go:generate'd
func (r Aws_Ec2_Subnet) Validate(t Template, context []string) (bool, []Failure) {
	var errors = make([]Failure, 0, 20)

	if ok, errs := r.AvailabilityZone.Validate(t, append(context, "AvailabilityZone")); !ok {
		errors = append(errors, errs...)
	}

	if ok, errs := r.CidrBlock.Validate(t, append(context, "CidrBlock")); !ok {
		errors = append(errors, errs...)
	}

	if ok, errs := r.MapPublicIpOnLaunch.Validate(t, append(context, "MapPublicIpOnLaunch")); !ok {
		errors = append(errors, errs...)
	}

	if ok, errs := r.Tags.Validate(t, append(context, "Tags")); !ok {
		errors = append(errors, errs...)
	}

	if ok, errs := r.VpcId.Validate(t, append(context, "VpcId")); !ok {
		errors = append(errors, errs...)
	}

	return len(errors) == 0, errors
}
