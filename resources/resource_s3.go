package resources

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/resources/common"
	. "github.com/jagregory/cfval/schema"
)

var s3LifecycleRule = NestedResource{
	Description: "AWS::S3::LifecycleRule",
	Properties: Properties{
		// "ExpirationDate":   Schema{Type: TypeString},
		"ExpirationInDays": Schema{
			Type: ValueNumber,
		},

		"Id": Schema{
			Type: ValueString,
		},

		// "NoncurrentVersionExpirationInDays": Schema{Type: TypeInteger},
		// "NoncurrentVersionTransition":       S3LifecycleRuleNoncurrentVersionTransition,
		// "Prefix":                            Schema{Type: TypeString},

		"Status": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		// "Transition":                        S3LifecycleRuleTransition,
	},
}

var accessControl = EnumValue{
	Description: "S3 Bucket AccessControl",

	Options: []string{"AuthenticatedRead", "AwsExecRead", "BucketOwnerRead", "BucketOwnerFullControl", "LogDeliveryWrite", "Private", "PublicRead", "PublicReadWrite"},
}

func Bucket() Resource {
	return Resource{
		AwsType: "AWS::S3::Bucket",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"AccessControl": Schema{
				Type: accessControl,
			},

			"BucketName": Schema{
				Type: ValueString,
			},

			// "CorsConfiguration":         s3_cors_configuration,

			"LifecycleConfiguration": Schema{
				Type: NestedResource{
					Description: "S3 Lifecycle Configuration",
					Properties: Properties{
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
				Type:  common.ResourceTag,
				Array: true,
			},

			// "VersioningConfiguration": s3_versioning_configuration,

			"WebsiteConfiguration": Schema{
				Type: NestedResource{
					Description: "S3 Website Configuration",
					Properties: Properties{
						"ErrorDocument": Schema{
							Type: ValueString,
						},

						"IndexDocument": Schema{
							Type:     ValueString,
							Required: constraints.Always,
						},

						// "RedirectAllRequestsTo": Schema{Type: ... }
						// "RoutingRules": Schema{Type: ...}
					},
				},
			},
		},
	}
}
