package resources

import . "github.com/jagregory/cfval/schema"

func LoadBalancer() Resource {
	return Resource{
		AwsType: "AWS::ElasticLoadBalancing::LoadBalancer",
		Properties: map[string]Schema{
			// AccessLoggingPolicy
			// Type: Elastic Load Balancing AccessLoggingPolicy

			// AppCookieStickinessPolicy
			// Type: A list of AppCookieStickinessPolicy objects.

			"AvailabilityZones": Schema{Type: TypeString, Array: true},

			"ConnectionDrainingPolicy": Schema{
				Type: Resource{
					AwsType: "Elastic Load Balancing ConnectionDrainingPolicy",
					Properties: map[string]Schema{
						"Enabled": Schema{Type: TypeBool, Required: true},
						"Timeout": Schema{Type: TypeInteger},
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
				Type: Resource{
					AwsType: "ElasticLoadBalancing HealthCheck",
					Properties: map[string]Schema{
						"HealthyThreshold":   Schema{Type: TypeString, Required: true},
						"Interval":           Schema{Type: TypeString, Required: true},
						"Target":             Schema{Type: TypeString, Required: true}, // TODO: Could be smarter about this restriction: "The protocol can be TCP, HTTP, HTTPS, or SSL. The range of valid ports is 1 through 65535."
						"Timeout":            Schema{Type: TypeString, Required: true}, // TODO: Could be smarter about this restriction: "This value must be less than the value for Interval."
						"UnhealthyThreshold": Schema{Type: TypeString, Required: true},
					},
				},
			},

			"Instances": ArrayOf(Schema{Type: TypeString}),

			// LBCookieStickinessPolicy
			// Type: A list of LBCookieStickinessPolicy objects.
			//
			// LoadBalancerName
			// Type: String

			"Listeners": Required(ArrayOf(Schema{
				Type: Resource{
					AwsType: "ElasticLoadBalancing Listener",
					Properties: map[string]Schema{
						"InstancePort":     Schema{Type: TypeString, Required: true},
						"InstanceProtocol": EnumOf("HTTP", "HTTPS", "TCP", "SSL"),
						"LoadBalancerPort": Schema{Type: TypeString, Required: true},
						"PolicyNames":      Schema{Type: TypeString, Array: true},
						"Protocol":         Required(EnumOf("HTTP", "HTTPS", "TCP", "SSL")),
						"SSLCertificateId": Schema{Type: TypeString},
					},
				},
			})),

			// Policies
			// Type: A list of ElasticLoadBalancing policy objects.
			//
			"Scheme": Schema{Type: TypeString},

			"SecurityGroups": Schema{Type: TypeString, Array: true},

			"Subnets": Schema{Type: TypeString, Array: true},

			"Tags": Schema{Type: resourceTag, Array: true},
		},
	}
}
