package schema

import "github.com/jagregory/cfval/reporting"

type Output struct {
	Description, Value interface{}
}

var outputSchema = Schema{
	Type:     ValueString,
	Required: Always,
}

func (o Output) Validate(template *Template, context []string) (reporting.ValidateResult, reporting.Failures) {
	if o.Description != nil {
		if _, ok := o.Description.(string); !ok {
			return reporting.ValidateOK, reporting.Failures{reporting.NewFailure("Expected a string", append(context, "Description"))}
		}
	}

	if _, errs := outputSchema.Validate(o.Value, TemplateResource{template: template}, append(context, "Value")); errs != nil {
		return reporting.ValidateOK, errs
	}

	return reporting.ValidateOK, nil
}
