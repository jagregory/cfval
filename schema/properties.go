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
		if value != nil && schema.Conflicts != nil && schema.Conflicts.Pass(values) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflict: %s", schema.Conflicts.Describe(values)), append(context, key)))
		}

		// TODO: can we merge Required, RequiredIf and RequiredUnless now we have
		// the richer Constraints

		// Validate RequiredIf
		if value == nil && schema.RequiredIf != nil && schema.RequiredIf.Pass(values) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Property is required: %s", schema.RequiredIf.Describe(values)), append(context, key)))
		}

		// Validate RequiredUnless
		// If this property isn't set AND any RequiredUnless properties are also
		// not set then we should fail because this property should be set
		if value == nil && schema.RequiredUnless != nil && !schema.RequiredUnless.Pass(values) {
			// TODO: this description is nasty
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Property is required: not %s", schema.RequiredUnless.Describe(values)), append(context, key)))
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
