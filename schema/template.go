package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type TemplateResource struct {
	template   *Template
	Definition Resource
	Properties map[string]interface{}
	Metadata   map[string]interface{}
}

func (tr TemplateResource) Template() *Template {
	return tr.template
}

func (tr TemplateResource) Property(name string) (interface{}, bool) {
	val, ok := tr.Properties[name]

	if ok {
		return val, ok
	}

	if def := tr.Definition.Properties[name]; def.Default != nil {
		return def.Default, true
	}

	return nil, false
}

func NewTemplateResource(template *Template) TemplateResource {
	return TemplateResource{template: template}
}

func (tr TemplateResource) Validate(context []string) (reporting.ValidateResult, reporting.Reports) {
	failures := make(reporting.Reports, 0, 50)

	if _, errs := tr.Definition.Validate(tr, context); errs != nil {
		failures = append(failures, errs...)
	}

	if _, errs := JSON.Validate(Schema{Type: JSON}, tr.Metadata, tr, append(context, "Metadata")); errs != nil {
		failures = append(failures, errs...)
	}

	if len(failures) == 0 {
		return reporting.ValidateOK, nil
	}

	return reporting.ValidateOK, failures
}

func (tr TemplateResource) HasProperty(name string, expected interface{}) bool {
	if value, found := tr.Properties[name]; found {
		return value == expected
	}

	return false
}

func NewUnrecognisedResource(template *Template, awsType string) TemplateResource {
	return TemplateResource{
		template: template,
		Definition: Resource{
			ValidateFunc: func(tr TemplateResource, context []string) (reporting.ValidateResult, reporting.Reports) {
				return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(fmt.Sprintf("Unrecognised resource %s", awsType), context)}
			},
		},
	}
}

type TemplateNestedResource struct {
	template *Template
	NestedResource
	Properties map[string]interface{}
}

func (tr TemplateNestedResource) Property(name string) (interface{}, bool) {
	val, ok := tr.Properties[name]
	return val, ok
}

func (r TemplateNestedResource) Template() *Template {
	return r.template
}

type Template struct {
	Resources  map[string]TemplateResource
	Parameters map[string]Parameter
	Outputs    map[string]Output
}

func (t *Template) Validate() (bool, reporting.Reports) {
	failures := make(reporting.Reports, 0, 100)

	for logicalID, resource := range t.Resources {
		if _, errs := resource.Validate([]string{"Resources", logicalID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for parameterID, parameter := range t.Parameters {
		if _, errs := parameter.Validate([]string{"Parameters", parameterID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	for outputID, output := range t.Outputs {
		if _, errs := output.Validate(t, []string{"Outputs", outputID}); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return true, nil
	}

	return false, failures
}
