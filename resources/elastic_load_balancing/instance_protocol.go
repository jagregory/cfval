package elastic_load_balancing

import . "github.com/jagregory/cfval/schema"

var instanceProtocol = EnumValue{
	Description: "LoadBalancer InstanceProtocol",

	Options: []string{"HTTP", "HTTPS", "TCP", "SSL"},
}
