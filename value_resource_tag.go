package main

var ResourceTag = Schema{
	Type: Resource{
		Properties: map[string]Schema{
			"Key": Schema{
				Type:     TypeString,
				Required: true,
			},

			"Value": Schema{
				Type:     TypeString,
				Required: true,
			},
		},
	},
}
