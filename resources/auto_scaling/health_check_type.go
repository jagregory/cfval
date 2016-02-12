package auto_scaling

import . "github.com/jagregory/cfval/schema"

var healthCheckType = EnumValue{
	Description: "Auto Scaling Health Check Type",

	Options: []string{"EC2", "ELB"},
}
