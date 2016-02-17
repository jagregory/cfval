package efs

import . "github.com/jagregory/cfval/schema"

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
func FileSystem() Resource {
	return Resource{
		AwsType: "AWS::EFS::FileSystem",

		// ID
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"FileSystemTags": Schema{
				Type:  fileSystemTag,
				Array: true,
			},
		},
	}
}
