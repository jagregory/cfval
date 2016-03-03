package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
var Table = Resource{
	AwsType: "AWS::DynamoDB::Table",

	Attributes: map[string]Schema{
		"StreamArn": Schema{
			Type: ARN,
		},
	},

	// Name
	ReturnValue: Schema{
		Type: ValueString,
	},

	Properties: Properties{
		"AttributeDefinitions": Schema{
			Type:     Multiple(attributeDefinition),
			Required: constraints.Always,
		},

		"GlobalSecondaryIndexes": Schema{
			Type: Multiple(globalSecondaryIndex),
		},

		"KeySchema": Schema{
			Type:     Multiple(keySchema),
			Required: constraints.Always,
		},

		"LocalSecondaryIndexes": Schema{
			Type: Multiple(localSecondaryIndex),
		},

		"StreamSpecification": Schema{
			Type: streamSpecification,
		},

		"ProvisionedThroughput": Schema{
			Type:     provisionedThroughput,
			Required: constraints.Always,
		},

		"TableName": Schema{
			Type: ValueString,
		},
	},
}
