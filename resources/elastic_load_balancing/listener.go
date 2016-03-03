package elastic_load_balancing

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
var listener = NestedResource{
	Description: "ElasticLoadBalancing Listener",
	Properties: Properties{
		"InstancePort": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"InstanceProtocol": Schema{
			Type: instanceProtocol,
			// TODO:
			// * If the front-end protocol is HTTP or HTTPS, InstanceProtocol has to
			//   be at the same protocol layer, i.e., HTTP or HTTPS. Likewise, if the
			//   front-end protocol is TCP or SSL, InstanceProtocol has to be TCP
			//   or SSL.
			// * If there is another listener with the same InstancePort whose
			//   InstanceProtocol is secure, i.e., HTTPS or SSL, the listener's
			//   InstanceProtocol has to be secure, i.e., HTTPS or SSL. If there is
			//   another listener with the same InstancePort whose InstanceProtocol is
			//   HTTP or TCP, the listener's InstanceProtocol must be either HTTP
			//   or TCP.
		},

		"LoadBalancerPort": Schema{
			Type:     ValueString,
			Required: constraints.Always,
		},

		"PolicyNames": Schema{
			Type: Multiple(ValueString),
		},

		"Protocol": Schema{
			Required: constraints.Always,
			Type:     instanceProtocol,
		},

		"SSLCertificateId": Schema{
			Type: ValueString,
		},
	},
}
