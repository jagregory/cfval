package resources

import . "github.com/jagregory/cfval/schema"

func Policy() Resource {
	return Resource{
		AwsType: "AWS::IAM::Policy",
		Properties: map[string]Schema{
			"Groups":         ArrayOf(Schema{Type: TypeString}),
			"PolicyDocument": Required(json),
			"PolicyName":     Schema{Type: TypeString, Required: true},
			"Roles":          ArrayOf(Schema{Type: TypeString}),
			"Users":          ArrayOf(Schema{Type: TypeString}),
		},
	}
}

func Role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",
		Properties: map[string]Schema{
			"AssumeRolePolicyDocument": Required(json),
			"ManagedPolicyArns":        ArrayOf(Schema{Type: TypeString}),
			"Path":                     Schema{Type: TypeString},
			"Policies": ArrayOf(Schema{
				Type: Resource{
					AwsType: "IAM Role Policy",
					Properties: map[string]Schema{
						"PolicyDocument": Required(json),
						"PolicyName":     Schema{Type: TypeString, Required: true},
					},
				},
			}),
		},
	}
}

func InstanceProfile() Resource {
	return Resource{
		AwsType: "AWS::IAM::InstanceProfile",
		Properties: map[string]Schema{
			"Path":  Schema{Type: TypeString, Required: true},
			"Roles": Required(ArrayOf(Schema{Type: TypeString})),
		},
	}
}
