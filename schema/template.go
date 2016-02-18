package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func TemplateValidate(template *parse.Template, definitions ResourceDefinitions) (bool, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	ctx := Context{
		Path:        make([]string, 0, 25),
		Template:    template,
		Definitions: definitions,
	}

	for logicalID, resource := range template.Resources {
		if _, errs := resourceValidate(*resource, ctx.Push("Resources", logicalID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for parameterID, parameter := range template.Parameters {
		if _, errs := parameterValidate(parameter, ctx.Push("Parameters", parameterID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for outputID, output := range template.Outputs {
		if _, errs := outputValidate(output, ctx.Push("Outputs", outputID)); errs != nil {
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

	definition := ctx.Definitions.Lookup(tr.Type)
	if _, errs := definition.Validate(ResourceWithDefinition{tr, definition}, ctx); errs != nil {
		failures = append(failures, errs...)
	}

	if _, errs := JSON.Validate(Schema{Type: JSON}, tr.Metadata, ResourceWithDefinition{tr, definition}, ctx.Push("Metadata")); errs != nil {
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

func outputValidate(o parse.Output, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	if o.Description != nil {
		if _, ok := o.Description.(string); !ok {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Expected a string", ctx.Push("Description").Path)}
		}
	}

	if _, errs := outputSchema.Validate(o.Value, emptyCurrentResource{}, ctx.Push("Value")); errs != nil {
		return reporting.ValidateOK, errs
	}

	return reporting.ValidateOK, nil
}
