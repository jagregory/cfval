package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-lsi.html
var localSecondaryIndex = NestedResource{
	Description: "DynamoDB Local Secondary Index",

	Properties: Properties{
		"IndexName": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: StringLengthValidate(3, 255),
		},

		"KeySchema": Schema{
			Type:     Multiple(keySchema),
			Required: constraints.Always,
		},

		"Projection": Schema{
			Type:     projection,
			Required: constraints.Always,
		},
	},
}
