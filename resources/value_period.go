package resources

import (
	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var period = Schema{
	Type: TypeString,
	ValidateFunc: func(interface{}, Template, []string) (bool, []reporting.Failure) {
		// TODO: Period. The time over which the specified statistic is applied. You must specify a time in seconds that is also a multiple of 60.
		return true, nil
	},
}
