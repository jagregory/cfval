package s3

import "github.com/jagregory/cfval/schema"

var storageClass = schema.EnumValue{
	Description: "S3 Storage Class",
	Options:     []string{"STANDARD_IA", "GLACIER"},
}
