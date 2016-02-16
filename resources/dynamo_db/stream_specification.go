package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-streamspecification.html
var streamSpecification = NestedResource{
	Description: "DynamoDB Table StreamSpecification",

	Properties: Properties{
		"StreamViewType": Schema{
			Type: EnumValue{
				Description: "StreamViewType",
				Options:     []string{"KEYS_ONLY", "NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES"},
			},
			Required: constraints.Always,
		},
	},
}
