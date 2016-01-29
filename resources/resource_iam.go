package resources

import . "github.com/jagregory/cfval/schema"

func Policy() Resource {
	return Resource{
		AwsType: "AWS::IAM::Policy",
		Properties: map[string]Schema{
			"Groups": Schema{
				Type:  TypeString,
				Array: true,
			},

			"PolicyDocument": Required(json),

			"PolicyName": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Roles": Schema{
				Type:  TypeString,
				Array: true,
			},

			"Users": Schema{
				Type:  TypeString,
				Array: true,
			},
		},
	}
}

func Role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",
		Properties: map[string]Schema{
			"AssumeRolePolicyDocument": Required(json),

			"ManagedPolicyArns": Schema{
				Type:  TypeString,
				Array: true,
			},

			"Path": Schema{
				Type: TypeString,
			},

			"Policies": Schema{
				Array: true,
				Type: Resource{
					AwsType: "IAM Role Policy",
					Properties: map[string]Schema{
						"PolicyDocument": Required(json),

						"PolicyName": Schema{
							Type:     TypeString,
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
		Properties: map[string]Schema{
			"Path": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Roles": Schema{
				Type:     TypeString,
				Array:    true,
				Required: true,
			},
		},
	}
}
