package schema

import "github.com/jagregory/cfval/reporting"

type Output struct {
	Description, Value interface{}
}

var outputSchema = Schema{Type: TypeString, Required: true}

func (o Output) Validate(template *Template, context []string) (bool, []reporting.Failure) {
	failures := make([]reporting.Failure, 0, 10)

	if _, ok := o.Description.(string); !ok {
		failures = append(failures, reporting.NewFailure("Expected a string", append(context, "Description")))
	}

	if ok, errs := validateProperty(outputSchema, o.Value, TemplateResource{Template: template}, append(context, "Value")); !ok {
		failures = append(failures, errs...)
	}

	return len(failures) == 0, failures
}
