package resources

import . "github.com/jagregory/cfval/schema"

var availabilityZone = EnumValidate(
	"ap-northeast-1a",
	"ap-northeast-1b",
	"ap-northeast-1c",
	"ap-northeast-2a",
	"ap-northeast-2c",
	"ap-southeast-1a",
	"ap-southeast-1b",
	"ap-southeast-2a",
	"ap-southeast-2b",
	"eu-central-1a",
	"eu-central-1b",
	"eu-west-1a",
	"eu-west-1b",
	"eu-west-1c",
	"sa-east-1a",
	"sa-east-1b",
	"sa-east-1c",
	"us-east-1a",
	"us-east-1b",
	"us-east-1c",
	"us-east-1d",
	"us-east-1e",
	"us-west-1a",
	"us-west-1c",
	"us-west-2a",
	"us-west-2b",
	"us-west-2c",
)
