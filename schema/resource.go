package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

type Resource struct {
	AwsType      string
	Properties   map[string]Schema
	ValidateFunc func(Template, map[string]interface{}, []string) (bool, []reporting.Failure)
}

func (rd Resource) Validate(t Template, tr TemplateResource, properties map[string]interface{}, context []string) (bool, []reporting.Failure) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(t, properties, context)
	}

	failures := make([]reporting.Failure, 0, 30)
	visited := make(map[string]bool)

	for key, schema := range rd.Properties {
		visited[key] = true

		value, _ := properties[key]
		if ok, errs := schema.Validate(value, t, tr, append(context, key)); !ok {
			failures = append(failures, errs...)
		}

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil {
			for _, conflict := range schema.Conflicts {
				if _, found := properties[conflict]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflicting property '%s' also set", conflict), append(context, key)))
				}
			}
		}

		// Validate RequiredIf
		if value == nil && schema.RequiredIf != nil {
			for _, required := range schema.RequiredIf {
				if _, found := properties[required]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is also set", required), append(context, key)))
				}
			}
		}

		// Validate RequiredUnless
		// If this property isn't set AND any RequiredUnless properties are also
		// not set then we should fail because this property should be set
		if value == nil && schema.RequiredUnless != nil {
			for _, required := range schema.RequiredUnless {
				if _, found := properties[required]; !found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is not set", required), append(context, key)))
				}
			}
		}
	}

	for key := range properties {
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
