package efs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-filesystemtags.html
var fileSystemTag = NestedResource{
	Description: "Elastic File System FileSystem FileSystemTag",

	Properties: Properties{
		"Key": Schema{
			Type: ValueString,
			// TODO: Implement this regex without negative lookaheads :(
			// 			 Suggestion, add a Negate nested validate...
			// ValidateFunc: RegexpValidate(
			// 	`^(:?[]).{1,128}$`,
			// 	"You can specify a value that is from 1 to 128 Unicode characters in length, but you cannot use the prefix aws:.",
			// ),
		},

		"Value": Schema{
			Type:         ValueString,
			ValidateFunc: StringLengthValidate(0, 128),
		},
	},
}
