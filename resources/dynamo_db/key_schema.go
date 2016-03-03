package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/deprecations"
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

		"HashKeyElement": Schema{
			Deprecated: deprecations.Deprecated("The HashKeyElement should now be defined in the AttributeDefinitions property of the AWS::DynamoDB::Table resource, and the AttributeName property of the KeySchema used them together."),
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
