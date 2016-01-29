package resources

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var vpcId = Schema{
	Type: TypeString,
	ValidateFunc: func(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
		// TODO: VpcId
		return true, nil
	},
}
