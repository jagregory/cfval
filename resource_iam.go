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

var IamPolicy = Schema{
	Type: policy(),
}

var Json = Schema{
	Type: TypeMap,
}

func role() Resource {
	return Resource{
		AwsType: "AWS::IAM::Role",
		Properties: map[string]Schema{
			"AssumeRolePolicyDocument": Required(Json),
			"ManagedPolicyArns":        ArrayOf(Schema{Type: TypeString}),
			"Path":                     Schema{Type: TypeString},
			"Policies":                 ArrayOf(IamPolicy),
		},
	}
}
