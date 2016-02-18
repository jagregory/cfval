package schema

import (
	"fmt"

	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// TODO: better name for this. It's either a TemplateResource or a "NestedTemplateResource"
type SelfRepresentation interface {
	Template() *parse.Template
	Property(name string) (interface{}, bool)
}

type Properties map[string]Schema

func (p Properties) PropertyDefault(name string) interface{} {
	return p[name].Default
}

func (p Properties) Validate(self constraints.CurrentResource, template *parse.Template, definitions ResourceDefinitions, path []string) (reporting.Reports, map[string]bool) {
	failures := make(reporting.Reports, 0, len(p)*2)
	visited := make(map[string]bool)

	for key, schema := range p {
		visited[key] = true
		value, _ := self.PropertyValue(key)

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil && schema.Conflicts.Pass(self) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Conflict: %s", schema.Conflicts.Describe(self)), append(path, key)))
		}

		// Validate Required
		if value == nil && schema.Required != nil && schema.Required.Pass(self) {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("%s is required because %s", key, schema.Required.Describe(self)), append(path, key)))
		}

		// assuming the above either failed and logged some failures, or passed and
		// we can safely skip over a nil property
		if value == nil {
			continue
		}

		if _, errs := schema.Validate(value, self, template, definitions, append(path, key)); errs != nil {
			failures = append(failures, errs...)
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
