package constraints

import "testing"

type TestCurrentResource map[string]interface{}

func (c TestCurrentResource) Properties() []string {
	props := make([]string, 0, len(c))
	for key := range c {
		props = append(props, key)
	}
	return props
}

func (c TestCurrentResource) PropertyDefault(string) interface{} {
	return nil
}

func (c TestCurrentResource) PropertyValue(name string) (interface{}, bool) {
	v, ok := c[name]
	return v, ok
}

func TestConstraintsPropertyExists(t *testing.T) {
	constraint := PropertyExists("MyProp")

	if !constraint.Pass(TestCurrentResource{"MyProp": "Value"}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(TestCurrentResource{}) {
		t.Error("Should fail if property doesn't exist")
	}
}
