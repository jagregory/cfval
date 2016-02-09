package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

// TODO: better name for this. It's either a TemplateResource or a "NestedTemplateResource"
type SelfRepresentation interface {
	Template() *Template
	Property(name string) (interface{}, bool)
}

type Properties map[string]Schema

func (p Properties) Validate(self SelfRepresentation, values map[string]interface{}, context []string) (reporting.Failures, map[string]bool) {
	failures := make(reporting.Failures, 0, len(p)*2)
	visited := make(map[string]bool)

	for key, schema := range p {
		visited[key] = true

		value, _ := values[key]
		if _, errs := schema.Validate(value, self, append(context, key)); errs != nil {
			failures = append(failures, errs...)
		}

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil {
			for _, conflict := range schema.Conflicts {
				if _, found := values[conflict]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflicting property '%s' also set", conflict), append(context, key)))
				}
			}
		}

		// Validate RequiredIf
		if value == nil && schema.RequiredIf != nil {
			for _, required := range schema.RequiredIf {
				if _, found := values[required]; found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is also set", required), append(context, key)))
				}
			}
		}

		// Validate RequiredUnless
		// If this property isn't set AND any RequiredUnless properties are also
		// not set then we should fail because this property should be set
		if value == nil && schema.RequiredUnless != nil {
			for _, required := range schema.RequiredUnless {
				if _, found := values[required]; !found {
					failures = append(failures, reporting.NewFailure(fmt.Sprintf("This property is required because '%s' is not set", required), append(context, key)))
				}
			}
		}
	}

	return failures, visited
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
