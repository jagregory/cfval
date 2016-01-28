package main

var S3LifecycleRule = Schema{
	Type: Resource{
		AwsType: "AWS::S3::LifecycleRule",
		Properties: map[string]Schema{
			// "ExpirationDate":   Schema{Type: TypeString},
			"ExpirationInDays": Schema{Type: TypeInteger},
			"Id":               Schema{Type: TypeString},
			// "NoncurrentVersionExpirationInDays": Schema{Type: TypeInteger},
			// "NoncurrentVersionTransition":       S3LifecycleRuleNoncurrentVersionTransition,
			// "Prefix":                            Schema{Type: TypeString},
			"Status": Schema{Type: TypeString, Required: true},
			// "Transition":                        S3LifecycleRuleTransition,
		},
	},
}

var S3LifecycleConfiguration = Schema{
	Type: Resource{
		AwsType: "AWS::S3::LifecycleConfiguration",
		Properties: map[string]Schema{
			"Rules": Required(ArrayOf(S3LifecycleRule)),
		},
	},
}

func bucket() Resource {
	return Resource{
		AwsType: "AWS::S3::Bucket",
		Properties: map[string]Schema{
			// "AccessControl":             EnumSchema("AuthenticatedRead", "AwsExecRead", "BucketOwnerRead", "BucketOwnerFullControl", "LogDeliveryWrite", "Private", "PublicRead", "PublicReadWrite"),
			"BucketName": Schema{Type: TypeString},
			// "CorsConfiguration":         s3_cors_configuration,
			"LifecycleConfiguration": S3LifecycleConfiguration,
			// "LoggingConfiguration":      s3_logging_configuration,
			// "NotificationConfiguration": s3_notification_configuration,
			// "ReplicationConfiguration":  s3_replication_configuration,
			"Tags": ArrayOf(ResourceTag),
			// "VersioningConfiguration": s3_versioning_configuration,
			// "WebsiteConfiguration":    s3_website_configuration,
		},
	}
}
