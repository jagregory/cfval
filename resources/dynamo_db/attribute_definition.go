package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-attributedef.html
var attributeDefinition = NestedResource{
	Description: "DynamoDB Attribute Definition",

	Properties: Properties{
		"AttributeName": Schema{
			Type:         ValueString,
			Required:     constraints.Always,
			ValidateFunc: StringLengthValidate(1, 255),
		},

		"AttributeType": Schema{
			Type: EnumValue{
				Description: "AttributeType",
				Options:     []string{"S", "N", "B"},
			},
			Required: constraints.Always,
		},
	},
}
