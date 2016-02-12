package route_53

import . "github.com/jagregory/cfval/schema"

var recordSetType = EnumValue{
	Description: "RecordSet Type",

	Options: []string{"A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SPF", "SRV", "TXT"},
}
