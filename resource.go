package main

type Resource struct {
	Properties   map[string]Schema
	ValidateFunc func(Template, map[string]interface{}, []string) (bool, []Failure)
}

func (rd Resource) Validate(t Template, properties map[string]interface{}, context []string) (bool, []Failure) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(t, properties, context)
	}

	failures := make([]Failure, 0, 30)
	for key, schema := range rd.Properties {
		value, _ := properties[key]
		if ok, errs := schema.Validate(value, t, append(context, key)); !ok {
			failures = append(failures, errs...)
		}
	}

	return len(failures) == 0, failures
}
