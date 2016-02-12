package cloud_front

import (
	"github.com/jagregory/cfval/constraints"
	. "github.com/jagregory/cfval/schema"
)

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
