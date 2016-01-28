package main

var Cidr = Schema{
	Type: TypeString,
	ValidateFunc: func(value interface{}, t Template, context []string) (bool, []Failure) {
		return true, nil
	},
}
