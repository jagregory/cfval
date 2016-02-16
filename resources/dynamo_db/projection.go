package dynamo_db

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-projectionobject.html
var projection = NestedResource{
	Description: "DynamoDB Projection",

	Properties: Properties{
		"NonKeyAttributes": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ProjectionType": Schema{
			Type: EnumValue{
				Description: "ProjectionType",
				Options:     []string{"KEYS_ONLY", "INCLUDE", "ALL"},
			},
		},
	},
}
