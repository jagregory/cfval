package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
func Table() Resource {
	return Resource{
		AwsType: "AWS::DynamoDB::Table",

		Attributes: map[string]Schema{
			"StreamArn": Schema{
				Type: ValueString,
			},
		},

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AttributeDefinitions": Schema{
				Type:     attributeDefinition,
				Array:    true,
				Required: constraints.Always,
			},

			"GlobalSecondaryIndexes": Schema{
				Type:  globalSecondaryIndex,
				Array: true,
			},

			"KeySchema": Schema{
				Type:     keySchema,
				Required: constraints.Always,
			},

			"LocalSecondaryIndexes": Schema{
				Type:  localSecondaryIndex,
				Array: true,
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
}
