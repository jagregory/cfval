package main

func policy() Resource {
	return Resource{
		AwsType: "AWS::IAM::Policy",
		Properties: map[string]Schema{
			"Groups":         ArrayOf(Schema{Type: TypeString}),
			"PolicyDocument": Required(Json),
			"PolicyName":     Schema{Type: TypeString, Required: true},
			"Roles":          ArrayOf(Schema{Type: TypeString}),
			"Users":          ArrayOf(Schema{Type: TypeString}),
		},
	}
}

func role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",
		Properties: map[string]Schema{
			"AssumeRolePolicyDocument": Required(Json),
			"ManagedPolicyArns":        ArrayOf(Schema{Type: TypeString}),
			"Path":                     Schema{Type: TypeString},
			"Policies": ArrayOf(Schema{
				Type: Resource{
					AwsType: "IAM Role Policy",
					Properties: map[string]Schema{
						"PolicyDocument": Required(Json),
						"PolicyName":     Schema{Type: TypeString, Required: true},
					},
				},
			}),
		},
	}
}

func instanceProfile() Resource {
	return Resource{
		AwsType: "AWS::IAM::InstanceProfile",
		Properties: map[string]Schema{
			"Path":  Schema{Type: TypeString, Required: true},
			"Roles": Required(ArrayOf(Schema{Type: TypeString})),
		},
	}
}
