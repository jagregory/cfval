package resources

import (
	"fmt"
	"regexp"

	"github.com/jagregory/cfval/reporting"
	. "github.com/jagregory/cfval/schema"
)

const cidrPattern = `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$`

var cidr = func(value interface{}, t Template, context []string) (bool, []reporting.Failure) {
	if ok, _ := regexp.MatchString(cidrPattern, value.(string)); !ok {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Cidr %s is invalid", value), context)}
	}

	return true, nil
}
