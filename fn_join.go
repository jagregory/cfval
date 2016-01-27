package main

type Join struct {
	Delim string
	Items []StringOrBuiltinFns
}

func (join Join) Validate(t Template, context []string) (bool, []Failure) {
	errors := make([]Failure, 0, len(join.Items))

	for _,item := range(join.Items) {
		if ok,errs := item.Validate(t, context); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
