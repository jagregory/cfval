package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

type Context struct {
	Template *parse.Template
	Path     []string
}

func (c Context) Push(path ...string) Context {
	return Context{
		Template: c.Template,
		Path:     append(c.Path, path...),
	}
}

func TemplateValidate(template *parse.Template, definitions ResourceDefinitions) (bool, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	ctx := Context{
		Template: template,
		Path:     make([]string, 0, 25),
	}

	for logicalID, resource := range template.Resources {
		if _, errs := resourceValidate(*resource, definitions, ctx.Push("Resources", logicalID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for parameterID, parameter := range template.Parameters {
		if _, errs := parameterValidate(parameter, ctx.Push("Parameters", parameterID)); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for outputID, output := range template.Outputs {
		if _, errs := outputValidate(output, definitions, ctx.Push("Outputs", outputID)); errs != nil {
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

func resourceValidate(tr parse.TemplateResource, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)

	definition := definitions.Lookup(tr.Type)
	if _, errs := definition.Validate(ResourceWithDefinition{tr, definition}, definitions, ctx); errs != nil {
		failures = append(failures, errs...)
	}

	if _, errs := JSON.Validate(Schema{Type: JSON}, tr.Metadata, ResourceWithDefinition{tr, definition}, definitions, ctx.Push("Metadata")); errs != nil {
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

func outputValidate(o parse.Output, definitions ResourceDefinitions, ctx Context) (reporting.ValidateResult, reporting.Reports) {
	if o.Description != nil {
		if _, ok := o.Description.(string); !ok {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure("Expected a string", ctx.Push("Description").Path)}
		}
	}

	if _, errs := outputSchema.Validate(o.Value, emptyCurrentResource{}, definitions, ctx.Push("Value")); errs != nil {
		return reporting.ValidateOK, errs
	}

	return reporting.ValidateOK, nil
}
