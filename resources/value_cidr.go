package resources

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var cidr = Schema{
	Type: TypeString,
	ValidateFunc: func(value interface{}, t Template, context []string) (bool, []reporting.Failure) {
		return true, nil
	},
}
