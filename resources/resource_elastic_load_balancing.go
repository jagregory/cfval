package resources

import . "github.com/jagregory/cfval/schema"

func LoadBalancer() Resource {
	return Resource{
		AwsType: "AWS::ElasticLoadBalancing::LoadBalancer",

		// Name
		ReturnValue: Schema{
			Type: ValueString,
		},

		Properties: Properties{
			// AccessLoggingPolicy
			// Type: Elastic Load Balancing AccessLoggingPolicy

			// AppCookieStickinessPolicy
			// Type: A list of AppCookieStickinessPolicy objects.

			"AvailabilityZones": Schema{
				Type:  ValueString,
				Array: true,
			},

			"ConnectionDrainingPolicy": Schema{
				Type: NestedResource{
					Description: "Elastic Load Balancing ConnectionDrainingPolicy",
					Properties: Properties{
						"Enabled": Schema{
							Type:     ValueBool,
							Required: true,
						},

						"Timeout": Schema{
							Type: ValueNumber,
						},
					},
				},
			},

			// Type: Elastic Load Balancing ConnectionDrainingPolicy
			//
			// ConnectionSettings
			// Type: Elastic Load Balancing ConnectionSettings
			//
			// CrossZone
			// Type: Boolean
			//
			"HealthCheck": Schema{
				// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
				Type: NestedResource{
					Description: "ElasticLoadBalancing HealthCheck",
					Properties: Properties{
						"HealthyThreshold": Schema{
							Type:     ValueString,
							Required: true,
						},

						"Interval": Schema{
							Type:     ValueString,
							Required: true,
						},

						"Target": Schema{
							Type:     ValueString,
							Required: true,
						}, // TODO: Could be smarter about this restriction: "The protocol can be TCP, HTTP, HTTPS, or SSL. The range of valid ports is 1 through 65535."

						"Timeout": Schema{
							Type:     ValueString,
							Required: true,
						}, // TODO: Could be smarter about this restriction: "This value must be less than the value for Interval."

						"UnhealthyThreshold": Schema{
							Type:     ValueString,
							Required: true,
						},
					},
				},
			},

			"Instances": Schema{
				Type:  ValueString,
				Array: true,
			},

			// LBCookieStickinessPolicy
			// Type: A list of LBCookieStickinessPolicy objects.
			//
			// LoadBalancerName
			// Type: String

			"Listeners": Schema{
				Array:    true,
				Required: true,
				Type: NestedResource{
					Description: "ElasticLoadBalancing Listener",
					Properties: Properties{
						"InstancePort": Schema{
							Type:     ValueString,
							Required: true,
						},

						"InstanceProtocol": Schema{
							Type: EnumValue{[]string{"HTTP", "HTTPS", "TCP", "SSL"}},
						},

						"LoadBalancerPort": Schema{
							Type:     ValueString,
							Required: true,
						},

						"PolicyNames": Schema{
							Type:  ValueString,
							Array: true,
						},

						"Protocol": Schema{
							Required: true,
							Type:     EnumValue{[]string{"HTTP", "HTTPS", "TCP", "SSL"}},
						},

						"SSLCertificateId": Schema{
							Type: ValueString,
						},
					},
				},
			},

			// Policies
			// Type: A list of ElasticLoadBalancing policy objects.
			//
			"Scheme": Schema{
				Type: ValueString,
			},

			"SecurityGroups": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Subnets": Schema{
				Type:  ValueString,
				Array: true,
			},

			"Tags": Schema{
				Type:  resourceTag,
				Array: true,
			},
		},
	}
}
