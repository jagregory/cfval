package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func TemplateValidate(template *parse.Template, definitions ResourceDefinitions, options ValidationOptions) (bool, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	ctx := NewInitialContext(template, definitions, options)

	for logicalID, resource := range template.Resources {
		if _, errs := resourceValidate(resource, ContextAdd(ctx, "Resources", logicalID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for parameterID, parameter := range template.Parameters {
		if _, errs := parameterValidate(parameter, ContextAdd(ctx, "Parameters", parameterID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for outputID, output := range template.Outputs {
		if _, errs := outputValidate(output, ContextAdd(ctx, "Outputs", outputID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return true, nil
	}

	return false, failures
}

func parameterValidate(p parse.Parameter, ctx Context) (bool, reporting.Reports) {
	// TODO: parameter validation?
	return true, nil
}

func resourceValidate(tr parse.TemplateResource, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)
	definition := ctx.Definitions().Lookup(tr.Type)

	resourceContext := NewResourceContext(ctx, ResourceWithDefinition{tr, definition})
	if _, errs := definition.Validate(resourceContext); errs != nil {
		failures = append(failures, errs...)
	}

	metadataContext := NewPropertyContext(resourceContext, Schema{Type: JSON})
	if _, errs := JSON.Validate(tr.Metadata, PropertyContextAdd(metadataContext, "Metadata")); errs != nil {
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

func outputValidate(o parse.Output, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	if o.Description != nil {
		if _, ok := o.Description.(string); !ok {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ContextAdd(ctx, "Description"), "Expected a string")}
		}
	}

	outputContext := NewResourceContext(ctx, emptyCurrentResource{})
	if _, errs := outputSchema.Validate(o.Value, ResourceContextAdd(outputContext, "Value")); errs != nil {
		return reporting.ValidateOK, errs
	}

	return reporting.ValidateOK, nil
}
