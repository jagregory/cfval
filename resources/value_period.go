package resources

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

var period = func(value interface{}, t Template, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	num, err := strconv.Atoi(value.(string))
	if err != nil {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Period is not a number: %s", value), context)}
	}

	if num == 0 || num%60 != 0 {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Period is not a multiple of 60: %s", value), context)}
	}

	return true, nil
}
