package schema

type emptyCurrentResource struct{}

func (emptyCurrentResource) PropertyValueOrDefault(name string) (interface{}, bool) {
	return nil, false
}

func (emptyCurrentResource) PropertyValue(string) (interface{}, bool) {
	return nil, false
}

func (emptyCurrentResource) PropertyDefault(string) (interface{}, bool) {
	return nil, false
}

func (emptyCurrentResource) Properties() []string {
	return []string{}
}
