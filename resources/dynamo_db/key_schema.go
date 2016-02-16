package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-keyschema.html
var keySchema = NestedResource{
	Description: "DynamoDB Key Schema",

	Properties: Properties{
		"AttributeName": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: StringLengthValidate(1, 255),
		},

		"KeyType": Schema{
			Type: EnumValue{
				Description: "KeyType",
				Options:     []string{"HASH", "RANGE"},
			},
			Required: constraints.Always,
		},
	},
}
