package main

type GetAtt struct {
	Attribute, Resource string
}

func (getatt GetAtt) Validate(t Template, context []string) (bool, []Failure) {
	// TODO
	return true, nil
}
