package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type Output struct {
	Description, Value interface{}
}

func (o Output) Validate(template *Template, context []string) (bool, []reporting.Failure) {
	failures := make([]reporting.Failure, 0, 10)

	if _, ok := o.Description.(string); !ok {
		failures = append(failures, reporting.NewFailure("Expected a string", append(context, "Description")))
	}

	fmt.Println("validating output", context)
	if ok, errs := validateProperty(Schema{Type: TypeString, Required: true}, o.Value, TemplateResource{Template: template}, append(context, "Value")); !ok {
		failures = append(failures, errs...)
	}

	return true, nil
}
