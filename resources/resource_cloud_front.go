package resources

import . "github.com/jagregory/cfval/schema"

var cookies = Resource{
	AwsType: "CloudFront ForwardedValues Cookies",
	Properties: map[string]Schema{
		"Forward": Schema{
			Type:     TypeString,
			Required: true,
		},

		"WhitelistedNames": Schema{
			Type:  TypeString,
			Array: true,
		},
	},
}

var forwardedValues = Resource{
	AwsType: "CloudFront ForwardedValues",
	Properties: map[string]Schema{
		"Cookies": Schema{
			Type: cookies,
		},

		"Headers": Schema{
			Type:  TypeString,
			Array: true,
		},

		"QueryString": Schema{
			Type:     TypeBool,
			Required: true,
		},
	},
}

var defaultCacheBehaviour = Resource{
	AwsType: "CloudFront DefaultCacheBehaviour",
	Properties: map[string]Schema{
		"AllowedMethods": Schema{
			Array: true,
			Type:  TypeString,
			ArrayValidateFunc: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
				[]string{"DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"},
			),
		},

		"CachedMethods": Schema{
			Array: true,
			Type:  TypeString,
			ArrayValidateFunc: FixedArrayValidate(
				[]string{"HEAD", "GET"},
				[]string{"GET", "HEAD", "OPTIONS"},
			),
		},

		"DefaultTTL": Schema{
			Type: TypeInteger,
		},

		"ForwardedValues": Schema{
			Required: true,
			Type:     forwardedValues,
		},

		"MaxTTL": Schema{
			Type: TypeInteger,
		},

		"MinTTL": Schema{
			Type: TypeString,
		},

		"SmoothStreaming": Schema{
			Type: TypeBool,
		},

		"TargetOriginId": Schema{
			Type:     TypeString,
			Required: true,
		},

		"TrustedSigners": Schema{
			Type:  TypeString,
			Array: true,
		},

		"ViewerProtocolPolicy": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

var logging = Resource{
	AwsType: "CloudFront Logging",
	Properties: map[string]Schema{
		"Bucket": Schema{
			Type:     TypeString,
			Required: true,
		},

		"IncludeCookies": Schema{
			Type: TypeBool,
		},

		"Prefix": Schema{
			Type: TypeString,
		},
	},
}

var customOriginConfig = Resource{
	AwsType: "CloudFront DistributionConfig Origin CustomOrigin",
	Properties: map[string]Schema{
		"HTTPPort": Schema{
			Type: TypeString,
		},

		"HTTPSPort": Schema{
			Type: TypeString,
		},

		"OriginProtocolPolicy": Schema{
			Type:     TypeString,
			Required: true,
		},
	},
}

var originConfig = Resource{
	AwsType: "CloudFront DistributionConfig Origin S3Origin",
	Properties: map[string]Schema{
		"OriginAccessIdentity": Schema{
			Type: TypeString,
		},
	},
}

var origin = Resource{
	AwsType: "CloudFront DistributionConfig Origin",
	Properties: map[string]Schema{
		"CustomOriginConfig": Schema{
			Type: customOriginConfig,
		},

		"DomainName": Schema{
			Type:     TypeString,
			Required: true,
		},

		"Id": Schema{
			Type:     TypeString,
			Required: true,
		},

		"OriginPath": Schema{
			Type: TypeString,
		},

		"S3OriginConfig": Schema{
			Type: originConfig,
		},
	},
}

var viewerCertificate = Resource{
	AwsType: "CloudFront DistributionConfiguration ViewerCertificate",
	Properties: map[string]Schema{
		"CloudFrontDefaultCertificate": Schema{
			Type: TypeBool,
		},

		"IamCertificateId": Schema{
			Type: TypeString,
		},

		"MinimumProtocolVersion": Schema{
			Type: TypeString,
		},

		"SslSupportMethod": Schema{
			Type: TypeString,
		},
	},
}

var distributionConfig = Resource{
	AwsType: "CloudFront DistributionConfig",
	Properties: map[string]Schema{
		"Aliases": Schema{
			Type:  TypeString,
			Array: true,
		},

		// "CacheBehaviors": ArrayOf(CacheBehavior),

		"Comment": Schema{
			Type: TypeString,
		},

		// "CustomErrorResponses": ArrayOf(CustomErrorResponse),

		"DefaultCacheBehavior": Schema{
			Required: true,
			Type:     defaultCacheBehaviour,
		},

		"DefaultRootObject": Schema{
			Type: TypeString,
		},

		"Enabled": Schema{
			Type:     TypeBool,
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
			Type: TypeString,
		},

		// "Restrictions": Restrictions,

		"ViewerCertificate": Schema{
			Type: viewerCertificate,
		},

		// "WebACLId": Schema{Type:TypeString},
	},
}

func Distribution() Resource {
	return Resource{
		AwsType: "AWS::CloudFront::Distribution",
		Properties: map[string]Schema{
			"DistributionConfig": Schema{
				Required: true,
				Type:     distributionConfig,
			},
		},
	}
}
