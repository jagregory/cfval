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

func (rd Resource) Validate(t Template, properties map[string]interface{}, context []string) (bool, []reporting.Failure) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(t, properties, context)
	}

	failures := make([]reporting.Failure, 0, 30)
	visited := make(map[string]bool)

	for key, schema := range rd.Properties {
		visited[key] = true

		value, _ := properties[key]
		if ok, errs := schema.Validate(value, t, append(context, key)); !ok {
			failures = append(failures, errs...)
		}
	}

	for key := range properties {
		if !visited[key] {
			failures = append(failures, reporting.NewFailure(fmt.Sprintf("Unknown property '%s' for %s", key, rd.AwsType), append(context, key)))
		}
	}

	return len(failures) == 0, failures
}
