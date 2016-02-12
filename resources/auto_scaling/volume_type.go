package auto_scaling

import . "github.com/jagregory/cfval/schema"

var volumeType = EnumValue{
	Description: "EBS Block Device Volume Type",

	Options: []string{"standard", "io1", "gp2"},
}
