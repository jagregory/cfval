package common

import . "github.com/jagregory/cfval/schema"

var EbsVolumeType = EnumValue{
	Description: "EBS Block Device Volume Type",

	Options: []string{"standard", "io1", "gp2"},
}
