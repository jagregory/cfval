package schema

import "github.com/jagregory/cfval/reporting"

// ValidatableProperties is a wrapper around Properties to allow us to support
// undefined or unsupported properties
type ValidatableProperties interface {
	PropertyDefault(string) (interface{}, bool)
	Validate(ResourceContext) reporting.Reports
	values() map[string]Schema
}

// UnsupportedProperties is a convinience type to give less intrusive errors for
// resource types we haven't implemented yet. People will get one "unsupported"
// warning per-resource, rather than lots of "unknown property" errors.
type UnsupportedProperties struct{}

func (UnsupportedProperties) PropertyDefault(name string) (interface{}, bool) {
	return nil, false
}

func (UnsupportedProperties) Validate(ctx ResourceContext) reporting.Reports {
	return reporting.Reports{
		reporting.NewWarning(ctx, "<unsupported> cfval does support validating %s resources yet", ctx.CurrentResource().AwsType()),
	}
}

func (UnsupportedProperties) values() map[string]Schema {
	return map[string]Schema{}
}

type Properties map[string]Schema

func (p Properties) PropertyDefault(name string) (interface{}, bool) {
	if def := p[name].Default; def != nil {
		return def, true
	}

	return nil, false
}

func (p Properties) Validate(ctx ResourceContext) reporting.Reports {
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

	// Reject any properties we weren't expecting
	for _, key := range self.Properties() {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(ResourceContextAdd(ctx, key), "%s is not a property of %s", key, self.AwsType()))
		}
	}

	return reporting.Safe(failures)
}

func (p Properties) values() map[string]Schema {
	return map[string]Schema(p)
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
