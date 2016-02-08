package resources

import . "github.com/jagregory/cfval/schema"

var cookies = NestedResource{
	Description: "CloudFront ForwardedValues Cookies",
	Properties: Properties{
		"Forward": Schema{
			Type:     ValueString,
			Required: true,
		},

		"WhitelistedNames": Schema{
			Type:  ValueString,
			Array: true,
		},
	},
}

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
			Required: true,
		},
	},
}

var defaultCacheBehaviour = NestedResource{
	Description: "CloudFront DefaultCacheBehaviour",
	Properties: Properties{
		"AllowedMethods": Schema{
			Type: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
				[]string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
			),
		},

		"CachedMethods": Schema{
			Type: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
			),
		},

		"DefaultTTL": Schema{
			Type: ValueNumber,
		},

		"ForwardedValues": Schema{
			Required: true,
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
			Required: true,
		},

		"TrustedSigners": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     ValueString,
			Required: true,
		},
	},
}

var logging = NestedResource{
	Description: "CloudFront Logging",
	Properties: Properties{
		"Bucket": Schema{
			Type:     ValueString,
			Required: true,
		},

		"IncludeCookies": Schema{
			Type: ValueBool,
		},

		"Prefix": Schema{
			Type: ValueString,
		},
	},
}

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
			Required: true,
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

var origin = NestedResource{
	Description: "CloudFront DistributionConfig Origin",
	Properties: Properties{
		"CustomOriginConfig": Schema{
			Type: customOriginConfig,
		},

		"DomainName": Schema{
			Type:     ValueString,
			Required: true,
		},

		"Id": Schema{
			Type:     ValueString,
			Required: true,
		},

		"OriginPath": Schema{
			Type: ValueString,
		},

		"S3OriginConfig": Schema{
			Type: originConfig,
		},
	},
}

var viewerCertificate = NestedResource{
	Description: "CloudFront DistributionConfiguration ViewerCertificate",
	Properties: Properties{
		"CloudFrontDefaultCertificate": Schema{
			Type: ValueBool,
		},

		"IamCertificateId": Schema{
			Type: ValueString,
		},

		"MinimumProtocolVersion": Schema{
			Type: ValueString,
		},

		"SslSupportMethod": Schema{
			Type: ValueString,
		},
	},
}

var cacheBehaviour = NestedResource{
	Description: "CloudFront DistributionConfig CacheBehavior",
	Properties: Properties{
		"AllowedMethods": Schema{
			Type: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
				[]string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
			),
		},

		"CachedMethods": Schema{
			Type: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
			),
		},

		"DefaultTTL": Schema{
			Type: ValueNumber,
		},

		"ForwardedValues": Schema{
			Type:     forwardedValues,
			Required: true,
		},

		"MaxTTL": Schema{
			Type: ValueNumber,
		},

		"MinTTL": Schema{
			Type: ValueNumber,
		},

		"PathPattern": Schema{
			Type:     ValueString,
			Required: true,
		},

		"SmoothStreaming": Schema{
			Type: ValueBool,
		},

		"TargetOriginId": Schema{
			Type:     ValueString,
			Required: true,
		},

		"TrustedSigners": Schema{
			Type:  ValueString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     EnumValue{[]string{"allow-all", "redirect-to-https", "https"}},
			Required: true,
		},
	},
}

var customErrorResponse = NestedResource{
	Description: "CloudFront DistributionConfig CustomErrorResponse",
	Properties:  Properties{},
}

var restrictions = NestedResource{
	Description: "CloudFront DistributionConfiguration Restrictions",
	Properties:  Properties{},
}

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
			Required: true,
			Type:     defaultCacheBehaviour,
		},

		"DefaultRootObject": Schema{
			Type: ValueString,
		},

		"Enabled": Schema{
			Type:     ValueBool,
			Required: true,
		},

		"Logging": Schema{
			Type: logging,
		},

		"Origins": Schema{
			Array:    true,
			Required: true,
			Type:     origin,
		},

		"PriceClass": Schema{
			Type: EnumValue{[]string{"PriceClass_All", "PriceClass_200", "PriceClass_100"}},
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
				Required: true,
				Type:     distributionConfig,
			},
		},
	}
}
