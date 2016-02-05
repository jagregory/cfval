package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Properties   Properties
	ReturnValue  Schema
	ValidateFunc func(TemplateResource, []string) (bool, []reporting.Failure)
}

func (rd Resource) Validate(tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(tr, context)
	}

	failures, visited := rd.Properties.Validate(tr, context)

	// Reject any properties we weren't expecting
	for key := range tr.Properties {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for %s", key, rd.AwsType), append(context, key)))
		}
	}

	return len(failures) == 0, failures
}

func collectKeys(m1 map[string]Schema, m2 map[string]interface{}) []string {
	set := make(map[string]bool)
	for key := range m1 {
		set[key] = true
	}
	for key := range m2 {
		set[key] = true
	}

	keys := make([]string, len(set))

	i := 0
	for k := range set {
		keys[i] = k
		i++
	}

	return keys
}

type Properties map[string]Schema

func (p Properties) Validate(tr TemplateResource, context []string) ([]reporting.Failure, map[string]bool) {
	failures := make([]reporting.Failure, 0, len(p)*2)
	visited := make(map[string]bool)

	for key, schema := range p {
		visited[key] = true

		value, _ := tr.Properties[key]
		if ok, errs := schema.Validate(value, tr, append(context, key)); !ok {
			failures = append(failures, errs...)
		}

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil {
			for _, conflict := range schema.Conflicts {
				if _, found := tr.Properties[conflict]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflicting property '%s' also set", conflict), append(context, key)))
				}
			}
		}

		// Validate RequiredIf
		if value == nil && schema.RequiredIf != nil {
			for _, required := range schema.RequiredIf {
				if _, found := tr.Properties[required]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is also set", required), append(context, key)))
				}
			}
		}

		// Validate RequiredUnless
		// If this property isn't set AND any RequiredUnless properties are also
		// not set then we should fail because this property should be set
		if value == nil && schema.RequiredUnless != nil {
			for _, required := range schema.RequiredUnless {
				if _, found := tr.Properties[required]; !found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is not set", required), append(context, key)))
				}
			}
		}
	}

	return failures, visited
}
