package constraints

import "testing"

func TestConstraintsPropertyExists(t *testing.T) {
	constraint := PropertyExists("MyProp")

	if !constraint.Pass(TestCurrentResource{map[string]interface{}{"MyProp": "Value"}, map[string]interface{}{}}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(TestCurrentResource{}) {
		t.Error("Should fail if property doesn't exist")
	}
}
