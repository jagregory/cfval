package resources

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues-cookies.html
var cookies = NestedResource{
	Description: "CloudFront ForwardedValues Cookies",

	Properties: Properties{
		"Forward": Schema{
			Type: EnumValue{
				Description: "CloudFront ForwardedValues Cookies Forward",
				Options:     []string{"none", "all", "whitelist"},
			},
			Required: constraints.Always,
		},

		"WhitelistedNames": Schema{
			Type:     ValueString,
			Array:    true,
			Required: constraints.PropertyIs("Forward", "whitelist"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-forwardedvalues.html
var forwardedValues = NestedResource{
	Description: "CloudFront ForwardedValues",
	Properties: Properties{
		"Cookies": Schema{
			Type: cookies,
		},

		"Headers": Schema{
			Type:  ValueString,
			Array: true,
		},

		"QueryString": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},
	},
}

var allowedMethods = FuncType{
	Description: "CloudFront Allowed Methods",

	Fn: FixedArrayValidate(
		[]string{"HEAD", "GET"},
		[]string{"GET", "HEAD", "OPTIONS"},
		[]string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
	),
}

var cachedMethods = FuncType{
	Description: "CloudFront Cached Methods",

	Fn: FixedArrayValidate(
		[]string{"HEAD", "GET"},
		[]string{"GET", "HEAD", "OPTIONS"},
	),
}

var defaultCacheBehaviour = NestedResource{
	Description: "CloudFront DefaultCacheBehaviour",
	Properties: Properties{
		"AllowedMethods": Schema{
			Type: allowedMethods,
		},

		"CachedMethods": Schema{
			Type: cachedMethods,
		},

		"DefaultTTL": Schema{
			Type: ValueNumber,
		},

		"ForwardedValues": Schema{
			Required: constraints.Always,
			Type:     forwardedValues,
		},

		"MaxTTL": Schema{
			Type: ValueNumber,
		},

		"MinTTL": Schema{
			Type: ValueString,
		},

		"SmoothStreaming": Schema{
			Type: ValueBool,
		},

		"TargetOriginId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TrustedSigners": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-logging.html
var logging = NestedResource{
	Description: "CloudFront Logging",
	Properties: Properties{
		"Bucket": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"IncludeCookies": Schema{
			Type: ValueBool,
		},

		"Prefix": Schema{
			Type: ValueString,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
var customOriginConfig = NestedResource{
	Description: "CloudFront DistributionConfig Origin CustomOrigin",
	Properties: Properties{
		"HTTPPort": Schema{
			Type: ValueString,
		},

		"HTTPSPort": Schema{
			Type: ValueString,
		},

		"OriginProtocolPolicy": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},
	},
}

var originConfig = NestedResource{
	Description: "CloudFront DistributionConfig Origin S3Origin",
	Properties: Properties{
		"OriginAccessIdentity": Schema{
			Type: ValueString,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
var origin = NestedResource{
	Description: "CloudFront DistributionConfig Origin",
	Properties: Properties{
		"CustomOriginConfig": Schema{
			Type:      customOriginConfig,
			Conflicts: constraints.PropertyExists("S3OriginConfig"),
			Required:  constraints.PropertyNotExists("S3OriginConfig"),
		},

		"DomainName": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"Id": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"OriginPath": Schema{
			Type:         ValueString,
			ValidateFunc: RegexpValidate(`^\/.*?[^\/]$`, "The value must start with a slash mark (/) and cannot end with a slash mark."),
		},

		"S3OriginConfig": Schema{
			Type:      originConfig,
			Conflicts: constraints.PropertyExists("CustomOriginConfig"),
			Required:  constraints.PropertyNotExists("CustomOriginConfig"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html
var viewerCertificate = NestedResource{
	Description: "CloudFront DistributionConfiguration ViewerCertificate",
	Properties: Properties{
		"CloudFrontDefaultCertificate": Schema{
			Type:      ValueBool,
			Conflicts: constraints.PropertyExists("IamCertificateId"),
			Required:  constraints.PropertyNotExists("IamCertificateId"),
		},

		"IamCertificateId": Schema{
			Type:      ValueString,
			Conflicts: constraints.PropertyExists("CloudFrontDefaultCertificate"),
			Required:  constraints.PropertyNotExists("CloudFrontDefaultCertificate"),
		},

		"MinimumProtocolVersion": Schema{
			Type: ValueString,
			// TODO: If you specify the IamCertificateId property and specify SNI only
			//       for the SslSupportMethod property, you must use TLSv1 for the
			//       minimum protocol version. If you don't specify a value, AWS
			//       CloudFormation specifies SSLv3.
		},

		"SslSupportMethod": Schema{
			Type:     ValueString,
			Required: constraints.PropertyExists("IamCertificateId"),
		},
	},
}

var viewerProtocolPolicy = EnumValue{
	Description: "CloudFront ViewerProtocolPolicy",

	Options: []string{"allow-all", "redirect-to-https", "https"},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-cachebehavior.html
var cacheBehaviour = NestedResource{
	Description: "CloudFront DistributionConfig CacheBehavior",
	Properties: Properties{
		"AllowedMethods": Schema{
			Type: allowedMethods,
		},

		"CachedMethods": Schema{
			Type: cachedMethods,
		},

		"DefaultTTL": Schema{
			Type: ValueNumber,
		},

		"ForwardedValues": Schema{
			Type:     forwardedValues,
			Required: constraints.Always,
		},

		"MaxTTL": Schema{
			Type: ValueNumber,
		},

		"MinTTL": Schema{
			Type: ValueNumber,
		},

		"PathPattern": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"SmoothStreaming": Schema{
			Type: ValueBool,
		},

		"TargetOriginId": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"TrustedSigners": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     viewerProtocolPolicy,
			Required: constraints.Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html
var customErrorResponse = NestedResource{
	Description: "CloudFront DistributionConfig CustomErrorResponse",
	Properties: Properties{
		"ErrorCachingMinTTL": Schema{
			Type: ValueNumber,
		},

		"ErrorCode": Schema{
			Type:         ValueNumber,
			Required:     constraints.Always,
			ValidateFunc: NumberOptions(400, 403, 404, 405, 414, 500, 501, 502, 503, 504),
		},

		"ResponseCode": Schema{
			Type:         ValueNumber,
			Required:     constraints.PropertyExists("ResponsePagePath"),
			ValidateFunc: NumberOptions(200, 400, 403, 404, 405, 414, 500, 501, 502, 503, 504),
		},

		"ResponsePagePath": Schema{
			Type:     ValueString,
			Required: constraints.PropertyExists("ResponseCode"),
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html
var geoRestriction = NestedResource{
	Description: "CloudFront DistributionConfig Restrictions GeoRestriction",
	Properties: Properties{
		"Locations": Schema{
			Type:  countryCode,
			Array: true,
			Required: constraints.Any{
				constraints.PropertyIs("RestrictionType", "blacklist"),
				constraints.PropertyIs("RestrictionType", "whitelist"),
			},
		},

		"RestrictionType": Schema{
			Type: EnumValue{
				Description: "RestrictionType",
				Options:     []string{"blacklist", "whitelist", "none"},
			},
			Required: constraints.Always,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
var restrictions = NestedResource{
	Description: "CloudFront DistributionConfiguration Restrictions",
	Properties: Properties{
		"GeoRestriction": Schema{
			Type:     geoRestriction,
			Required: constraints.Always,
		},
	},
}

var priceClass = EnumValue{
	Description: "CloudFront PriceClass",

	Options: []string{"PriceClass_All", "PriceClass_200", "PriceClass_100"},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig.html
var distributionConfig = NestedResource{
	Description: "CloudFront DistributionConfig",
	Properties: Properties{
		"Aliases": Schema{
			Type:  ValueString,
			Array: true,
		},

		"CacheBehaviors": Schema{
			Type:  cacheBehaviour,
			Array: true,
		},

		"Comment": Schema{
			Type: ValueString,
		},

		"CustomErrorResponses": Schema{
			Type:  customErrorResponse,
			Array: true,
		},

		"DefaultCacheBehavior": Schema{
			Required: constraints.Always,
			Type:     defaultCacheBehaviour,
		},

		"DefaultRootObject": Schema{
			Type: ValueString,
		},

		"Enabled": Schema{
			Type:     ValueBool,
			Required: constraints.Always,
		},

		"Logging": Schema{
			Type: logging,
		},

		"Origins": Schema{
			Array:    true,
			Required: constraints.Always,
			Type:     origin,
		},

		"PriceClass": Schema{
			Type: priceClass,
		},

		"Restrictions": Schema{
			Type: restrictions,
		},

		"ViewerCertificate": Schema{
			Type: viewerCertificate,
		},

		"WebACLId": Schema{
			Type: ValueString,
		},
	},
}

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html
func Distribution() Resource {
	return Resource{
		AwsType: "AWS::CloudFront::Distribution",

		// Distribution ID
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			"DistributionConfig": Schema{
				Required: constraints.Always,
				Type:     distributionConfig,
			},
		},
	}
}
