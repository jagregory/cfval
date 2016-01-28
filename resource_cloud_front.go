package main

func distribution() Resource {
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
									"AllowedMethods": ArrayOf(EnumSchema("DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT")),

									"CachedMethods": ArrayOf(EnumSchema("GET", "HEAD", "OPTIONS")),

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
								AwsType:    "CloudFront Distribution Logging",
								Properties: map[string]Schema{},
							},
						},
						"Origins": ArrayOf(Schema{
							Required: true,
							Type: Resource{
								AwsType:    "CloudFront Distribution Origin",
								Properties: map[string]Schema{},
							},
						}),
						"PriceClass": Schema{Type: TypeString},
						// "Restrictions": Restrictions,
						"ViewerCertificate": Schema{
							Type: Resource{
								AwsType:    "CloudFront Distribution ViewerCertificate",
								Properties: map[string]Schema{},
							},
						},
						// "WebACLId": Schema{Type:TypeString},
					},
				},
			}),
		},
	}
}
