package resources

import . "github.com/jagregory/cfval/schema"

var s3LifecycleRule = Resource{
	AwsType: "AWS::S3::LifecycleRule",
	Properties: map[string]Schema{
		// "ExpirationDate":   Schema{Type: TypeString},
		"ExpirationInDays": Schema{
			Type: TypeInteger,
		},

		"Id": Schema{
			Type: TypeString,
		},

		// "NoncurrentVersionExpirationInDays": Schema{Type: TypeInteger},
		// "NoncurrentVersionTransition":       S3LifecycleRuleNoncurrentVersionTransition,
		// "Prefix":                            Schema{Type: TypeString},

		"Status": Schema{
			Type:     TypeString,
			Required: true,
		},

		// "Transition":                        S3LifecycleRuleTransition,
	},
}

func Bucket() Resource {
	return Resource{
		AwsType: "AWS::S3::Bucket",

		// Name
		ReturnValue: Schema{
			Type: TypeString,
		},

		Properties: map[string]Schema{
			"AccessControl": Schema{
				Type:         TypeString,
				ValidateFunc: EnumValidate("AuthenticatedRead", "AwsExecRead", "BucketOwnerRead", "BucketOwnerFullControl", "LogDeliveryWrite", "Private", "PublicRead", "PublicReadWrite"),
			},

			"BucketName": Schema{
				Type: TypeString,
			},

			// "CorsConfiguration":         s3_cors_configuration,

			"LifecycleConfiguration": Schema{
				Type: Resource{
					AwsType: "S3 Lifecycle Configuration",
					Properties: map[string]Schema{
						"Rules": Schema{
							Type:  s3LifecycleRule,
							Array: true,
						},
					},
				},
			},

			// "LoggingConfiguration":      s3_logging_configuration,
			// "NotificationConfiguration": s3_notification_configuration,
			// "ReplicationConfiguration":  s3_replication_configuration,

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},

			// "VersioningConfiguration": s3_versioning_configuration,

			"WebsiteConfiguration": Schema{
				Type: Resource{
					AwsType: "S3 Website Configuration",
					Properties: map[string]Schema{
						"ErrorDocument": Schema{
							Type: TypeString,
						},

						"IndexDocument": Schema{
							Type:     TypeString,
							Required: true,
						},

						// "RedirectAllRequestsTo": Schema{Type: ... }
						// "RoutingRules": Schema{Type: ...}
					},
				},
			},
		},
	}
}
