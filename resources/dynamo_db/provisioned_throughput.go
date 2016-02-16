package dynamo_db

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html
var provisionedThroughput = NestedResource{
	Description: "DynamoDB Provisioned Throughput",

	Properties: Properties{
		"ReadCapacityUnits": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},

		"WriteCapacityUnits": Schema{
			Type:     ValueNumber,
			Required: constraints.Always,
		},
	},
}
