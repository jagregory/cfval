package resources

import . "github.com/jagregory/cfval/schema"

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
				Required: true,
			},

			"PolicyName": Schema{
				Type:     ValueString,
				Required: true,
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
				Required: true,
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
							Required: true,
						},

						"PolicyName": Schema{
							Type:     ValueString,
							Required: true,
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
				Required: true,
			},

			"Roles": Schema{
				Type:     ValueString,
				Array:    true,
				Required: true,
			},
		},
	}
}
