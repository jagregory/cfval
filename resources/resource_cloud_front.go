package resources

import . "github.com/jagregory/cfval/schema"

func Distribution() Resource {
	return Resource{
		AwsType: "AWS::CloudFront::Distribution",
		Properties: map[string]Schema{
			"DistributionConfig": Required(Schema{
				Type: Resource{
					AwsType: "CloudFront DistributionConfig",
					Properties: map[string]Schema{
						"Aliases": ArrayOf(Schema{Type: TypeString}),
						// "CacheBehaviors": ArrayOf(CacheBehavior),
						"Comment": Schema{Type: TypeString},
						// "CustomErrorResponses": ArrayOf(CustomErrorResponse),
						"DefaultCacheBehavior": Schema{
							Required: true,
							Type: Resource{
								AwsType: "CloudFront DefaultCacheBehaviour",
								Properties: map[string]Schema{
									"AllowedMethods": ArrayOf(EnumOf("DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT")),

									"CachedMethods": ArrayOf(EnumOf("GET", "HEAD", "OPTIONS")),

									"DefaultTTL": Schema{Type: TypeInteger},

									"ForwardedValues": Schema{
										Required: true,
										Type: Resource{
											AwsType: "CloudFront ForwardedValues",
											Properties: map[string]Schema{
												"Cookies": Schema{
													Type: Resource{
														AwsType: "CloudFront ForwardedValues Cookies",
														Properties: map[string]Schema{
															"Forward": Schema{Type: TypeString, Required: true},

															"WhitelistedNames": ArrayOf(Schema{Type: TypeString}),
														},
													},
												},

												"Headers": ArrayOf(Schema{Type: TypeString}),

												"QueryString": Schema{Type: TypeBool, Required: true},
											},
										},
									},

									"MaxTTL": Schema{Type: TypeInteger},

									"MinTTL": Schema{Type: TypeString},

									"SmoothStreaming": Schema{Type: TypeBool},

									"TargetOriginId": Schema{Type: TypeString, Required: true},

									"TrustedSigners": ArrayOf(Schema{Type: TypeString}),

									"ViewerProtocolPolicy": Schema{Type: TypeString, Required: true},
								},
							},
						},
						"DefaultRootObject": Schema{Type: TypeString},
						"Enabled":           Schema{Type: TypeBool, Required: true},
						"Logging": Schema{
							Type: Resource{
								AwsType: "CloudFront Logging",
								Properties: map[string]Schema{
									"Bucket":         Schema{Type: TypeString, Required: true},
									"IncludeCookies": Schema{Type: TypeBool},
									"Prefix":         Schema{Type: TypeString},
								},
							},
						},
						"Origins": ArrayOf(Schema{
							Required: true,
							Type: Resource{
								AwsType: "CloudFront DistributionConfig Origin",
								Properties: map[string]Schema{
									"CustomOriginConfig": Schema{
										Type: Resource{
											AwsType: "CloudFront DistributionConfig Origin CustomOrigin",
											Properties: map[string]Schema{
												"HTTPPort": Schema{Type: TypeString},

												"HTTPSPort": Schema{Type: TypeString},

												"OriginProtocolPolicy": Schema{Type: TypeString, Required: true},
											},
										},
									},

									"DomainName": Schema{Type: TypeString, Required: true},

									"Id": Schema{Type: TypeString, Required: true},

									"OriginPath": Schema{Type: TypeString},

									"S3OriginConfig": Schema{
										Type: Resource{
											AwsType: "CloudFront DistributionConfig Origin S3Origin",
											Properties: map[string]Schema{
												"OriginAccessIdentity": Schema{Type: TypeString},
											},
										},
									},
								},
							},
						}),

						"PriceClass": Schema{Type: TypeString},
						// "Restrictions": Restrictions,
						"ViewerCertificate": Schema{
							Type: Resource{
								AwsType: "CloudFront DistributionConfiguration ViewerCertificate",
								Properties: map[string]Schema{
									"CloudFrontDefaultCertificate": Schema{Type: TypeBool},
									"IamCertificateId":             Schema{Type: TypeString},
									"MinimumProtocolVersion":       Schema{Type: TypeString},
									"SslSupportMethod":             Schema{Type: TypeString},
								},
							},
						},
						// "WebACLId": Schema{Type:TypeString},
					},
				},
			}),
		},
	}
}
