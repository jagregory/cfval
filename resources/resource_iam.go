package resources

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

func Policy() Resource {
	return Resource{
		AwsType: "AWS::IAM::Policy",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"Groups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"PolicyDocument": Schema{
				Type:     JSON,
				Required: constraints.Always,
			},

			"PolicyName": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Roles": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Users": Schema{
				Type:  ValueString,
				Array: true,
			},
		},
	}
}

func Role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"AssumeRolePolicyDocument": Schema{
				Type:     JSON,
				Required: constraints.Always,
			},

			"ManagedPolicyArns": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Path": Schema{
				Type: ValueString,
			},

			"Policies": Schema{
				Array: true,
				Type: NestedResource{
					Description: "IAM Role Policy",
					Properties: map[string]Schema{
						"PolicyDocument": Schema{
							Type:     JSON,
							Required: constraints.Always,
						},

						"PolicyName": Schema{
							Type:     ValueString,
							Required: constraints.Always,
						},
					},
				},
			},
		},
	}
}

func InstanceProfile() Resource {
	return Resource{
		AwsType: "AWS::IAM::InstanceProfile",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: map[string]Schema{
			"Path": Schema{
				Type:     ValueString,
				Required: constraints.Always,
			},

			"Roles": Schema{
				Type:     ValueString,
				Array:    true,
				Required: constraints.Always,
			},
		},
	}
}
