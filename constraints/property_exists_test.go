package constraints

import "testing"

type TestCurrentResource struct {
	values   map[string]interface{}
	defaults map[string]interface{}
}

func (c TestCurrentResource) Properties() []string {
	props := make([]string, 0, len(c.values))
	for key := range c.values {
		props = append(props, key)
	}
	return props
}

func (c TestCurrentResource) PropertyValueOrDefault(name string) (interface{}, bool) {
	if v, ok := c.PropertyValue(name); ok {
		return v, true
	}

	if v, ok := c.PropertyDefault(name); ok {
		return v, true
	}

	return nil, false
}

func (c TestCurrentResource) PropertyDefault(name string) (interface{}, bool) {
	v, ok := c.defaults[name]
	return v, ok
}

func (c TestCurrentResource) PropertyValue(name string) (interface{}, bool) {
	v, ok := c.values[name]
	return v, ok
}

func TestConstraintsPropertyExists(t *testing.T) {
	constraint := PropertyExists("MyProp")

	if !constraint.Pass(TestCurrentResource{map[string]interface{}{"MyProp": "Value"}, map[string]interface{}{}}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(TestCurrentResource{}) {
		t.Error("Should fail if property doesn't exist")
	}
}
