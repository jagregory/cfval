package main

type Template struct {
	Resources  map[string]Resource
	Parameters map[string]Parameter
}

func (t Template) Validate() (bool, []Failure) {
	errors := make([]Failure, 0, 100)

	for logicalId, resource := range t.Resources {
		if ok, errs := resource.Validate(t, []string{"Resources", logicalId}); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
