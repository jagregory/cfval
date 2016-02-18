package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func TemplateValidate(t *parse.Template, definitions ResourceDefinitions) (bool, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	for logicalID, resource := range t.Resources {
		if _, errs := resourceValidate(*resource, definitions, []string{"Resources", logicalID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for parameterID, parameter := range t.Parameters {
		if _, errs := parameterValidate(parameter, []string{"Parameters", parameterID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for outputID, output := range t.Outputs {
		if _, errs := outputValidate(output, t, definitions, []string{"Outputs", outputID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return true, nil
	}

	return false, failures
}

func parameterValidate(p parse.Parameter, context []string) (bool, reporting.Reports) {
	// TODO: parameter validation?
	return true, nil
}

func resourceValidate(tr parse.TemplateResource, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)

	definition := definitions.Lookup(tr.Type)
	if _, errs := definition.Validate(ResourceWithDefinition{tr, definition}, tr.Template(), definitions, context); errs != nil {
		failures = append(failures, errs...)
	}

	if _, errs := JSON.Validate(Schema{Type: JSON}, tr.Metadata, ResourceWithDefinition{tr, definition}, tr.Tmpl, definitions, append(context, "Metadata")); errs != nil {
		failures = append(failures, errs...)
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

var outputSchema = Schema{
	Type:     ValueString,
	Required: constraints.Always,
}

type emptyCurrentResource struct {
}

func (emptyCurrentResource) PropertyValue(string) (interface{}, bool) {
	return nil, false
}

func (emptyCurrentResource) PropertyDefault(string) interface{} {
	return nil
}

func (emptyCurrentResource) Properties() []string {
	return []string{}
}

func outputValidate(o parse.Output, template *parse.Template, definitions ResourceDefinitions, context []string) (reporting.ValidateResult, reporting.Reports) {
	if o.Description != nil {
		if _, ok := o.Description.(string); !ok {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Expected a string", append(context, "Description"))}
		}
	}

	if _, errs := outputSchema.Validate(o.Value, emptyCurrentResource{}, template, definitions, append(context, "Value")); errs != nil {
		return reporting.ValidateOK, errs
	}

	return reporting.ValidateOK, nil
}
