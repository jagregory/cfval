package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// TODO: better name for this. It's either a TemplateResource or a "NestedTemplateResource"
type SelfRepresentation interface {
	Template() *parse.Template
	Property(name string) (interface{}, bool)
}

type Properties map[string]Schema

func (p Properties) PropertyDefault(name string) (interface{}, bool) {
	if def := p[name].Default; def != nil {
		return def, true
	}

	return nil, false
}

func (p Properties) Validate(ctx ResourceContext) (reporting.Reports, map[string]bool) {
	failures := make(reporting.Reports, 0, len(p)*2)
	visited := make(map[string]bool)

	self := ctx.CurrentResource()

	for key, schema := range p {
		visited[key] = true
		value, _ := self.PropertyValue(key)

		// Validate conflicting properties
		if value != nil && schema.Conflicts != nil && schema.Conflicts.Pass(self) {
			failures = append(failures, reporting.NewFailure(ResourceContextAdd(ctx, key), "Conflict: %s", schema.Conflicts.Describe(self)))
		}

		// Validate Required
		if value == nil && schema.Required != nil && schema.Required.Pass(self) {
			failures = append(failures, reporting.NewFailure(ResourceContextAdd(ctx, key), "%s is required because %s", key, schema.Required.Describe(self)))
		}

		// assuming the above either failed and logged some failures, or passed and
		// we can safely skip over a nil property
		if value == nil {
			continue
		}

		if _, errs := schema.Validate(value, ResourceContextAdd(ctx, key)); errs != nil {
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
